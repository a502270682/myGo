package route

import (
	"context"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/globalsign/mgo/bson"
	"github.com/pkg/errors"
	"myGo/adapter/error_code"
	"myGo/adapter/log"
	"net/http"
	"net/url"
	"reflect"
)

const defaultMemory = 32 * 1024 * 1024

type originFormKey struct{}
type bindFormKey struct{}

var (
	EmptyStr = ""

	ErrMustPtr           = errors.New("param must be ptr")
	ErrMustPointToStruct = errors.New("param must point to struct")
	ErrMustHasThreeParam = errors.New("method must has three input")
	ErrMustFunc          = errors.New("method must be func")
	ErrMustValid         = errors.New("method must be valid")
	ErrMustReplyErrorPtr = errors.New("method ret must be *error_code.ReplyError")
	ErrMustOneOut        = errors.New("method must has one out")

	initerType    = reflect.TypeOf((*Initer)(nil)).Elem()
	OriginFormKey = originFormKey{} //原始的form
	BindFormKey   = bindFormKey{}   //删除掉空值后的form
)

type Initer interface {
	Init(ctx context.Context)
}

func init() {
	structs.DefaultTagName = "json"
}

func Route(routes gin.IRoutes, method string, path string, function interface{}) {
	routes.Handle(method, path, CreateHandlerFunc(function))
}

func checkMethod(method interface{}) (mV reflect.Value, reqT, replyT reflect.Type, err error) {
	mV = reflect.ValueOf(method)
	if !mV.IsValid() {
		err = ErrMustValid
		return
	}
	mT := mV.Type()
	if mT.Kind() != reflect.Func {
		err = ErrMustFunc
		return
	}
	if mT.NumIn() != 3 {
		err = ErrMustHasThreeParam
		return
	}
	reqT = mT.In(1)
	if reqT.Kind() != reflect.Ptr {
		err = ErrMustPtr
		return
	}
	if reqT.Elem().Kind() != reflect.Struct {
		err = ErrMustPointToStruct
		return
	}
	reqT = reqT.Elem()
	replyT = mT.In(2)
	if replyT.Kind() != reflect.Ptr {
		err = ErrMustPtr
		return
	}
	if replyT.Elem().Kind() != reflect.Struct {
		err = ErrMustPointToStruct
		return
	}
	replyT = replyT.Elem()
	if mT.NumOut() != 1 {
		err = ErrMustOneOut
		return
	}
	//retT := mT.Out(0)
	//if retT != replyErrorType {
	//	err = ErrMustReplyErrorPtr
	//	return
	//}
	return mV, reqT, replyT, err
}

func isImplementIniter(v reflect.Value) bool {
	return v.Type().Implements(initerType)
}

func processReq(ctx context.Context, v reflect.Value) {
	elem := v.Elem()
	for i := 0; i < elem.NumField(); i++ {
		ev := elem.Field(i)
		if ev.CanInterface() {
			it := ev.Interface()
			if id, ok := it.(bson.ObjectId); ok {
				t := string(id)
				if bson.IsObjectIdHex(t) {
					nid := bson.ObjectIdHex(t)
					ev.SetString(string(nid))
				} else {
					ev.SetString("")
				}
			}
		}
	}
}

func callFieldInit(ctx context.Context, v reflect.Value) {
	elem := v.Elem()
	vT := elem.Type()
	for i := 0; i < elem.NumField(); i++ {
		ev := elem.Field(i)
		if isImplementIniter(ev) {
			if ev.CanSet() {
				ev.Set(reflect.New(vT.Field(i).Type.Elem()))
				initer := ev.Interface().(Initer)
				initer.Init(ctx)
			}
		}
	}
}

func structToMap(t interface{}) map[string]interface{} {
	return structs.Map(t)
}

func CreateHandlerFunc(method interface{}) gin.HandlerFunc {
	return CreateHandlerFuncWithLogger(method, log.GetLogger())
}

func CreateHandlerFuncWithLogger(method interface{}, l log.Logger) gin.HandlerFunc {
	mV, reqT, replyT, err := checkMethod(method)
	if err != nil {
		panic(err)
	}

	return func(c *gin.Context) {
		ctx := c.Request.Context()
		req := reflect.New(reqT)
		_ = c.Request.ParseForm()
		_ = c.Request.ParseMultipartForm(defaultMemory)

		formValue := c.Request.Form
		originForm := url.Values{}
		for k, vs := range formValue {
			originForm[k] = vs
			t := []string{}
			for _, v := range vs {
				if v != "" {
					t = append(t, v)
				}
			}
			if len(t) == 0 {
				delete(formValue, k)
			} else {
				formValue[k] = t
			}
		}
		log.WithField("form", c.Request.Form).Debug(ctx, "request param")
		if err := c.ShouldBindWith(req.Interface(), binding.Form); err != nil {
			replyError := error_code.Error(error_code.CODE_PARAM_WRONG, err.Error())
			c.JSON(http.StatusOK, replyError)
			l.WithFields(log.Fields{
				"req": c.Request.URL,
				"err": err,
			}).Warn(ctx, "bind param failed")
			return
		}

		c.Request.Form = originForm //把request里面的Form还原为原始form，以便后续使用
		ctx = context.WithValue(ctx, OriginFormKey, originForm)
		ctx = context.WithValue(ctx, BindFormKey, formValue)

		callFieldInit(ctx, req)
		processReq(ctx, req)

		reply := reflect.New(replyT)
		l.WithFields(log.Fields{
			"func": mV.Type().String(),
			"req":  req,
		}).Debugf(ctx, "invoke handler")
		results := mV.Call([]reflect.Value{reflect.ValueOf(ctx), req, reply})
		err, _ := results[0].Interface().(*error_code.ReplyError)
		if err != nil {
			if err.IsAutoMsg() {
				msg := error_code.ErrCodeMessage(err.Code)
				if msg != "" {
					err.Message = msg
				}
			}
			l.WithFields(log.Fields{
				"req": c.Request.URL,
				"err": err,
			}).Warn(ctx, "handler err")
			c.JSON(0, err)
			return
		}
		m := structToMap(reply.Interface())
		if _, ok := m["code"]; !ok {
			m["code"] = 0
		}
		if _, ok := m["message"]; !ok {
			m["message"] = ""
		}
		c.PureJSON(http.StatusOK, m)
	}
}
