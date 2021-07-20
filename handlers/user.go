package handlers

import (
	"context"
	"fmt"
	"myGo/adapter/error_code"
	"myGo/adapter/log"
	"myGo/models"
	"myGo/proto"
)

func UserInfoHandler(ctx context.Context, req *proto.UserInfoReq, rsp *proto.UserInfoRsp) *error_code.ReplyError {
	user, err := models.GetUserDao().GetUser(req.Name)
	if err != nil {
		log.Errorf(ctx, "fail to call GetUser, err:%+v", err)
		return error_code.Error(error_code.CodeParamWrong, "")
	}
	result := models.GetUserDao().GetUserWithSchool(req.Name)
	for _, t := range result {
		fmt.Println(t)
	}
	//schools, _ := models.GetSchoolDao().GetSchoolsByNames([]string{"绵阳中学", "南山中学"})
	//for _, t := range schools {
	//	fmt.Println(t)
	//}
	s, _ := models.GetSchoolDao().GetSchoolsByType([]models.SchoolType{models.SchoolTypeJSchool})
	for _, t := range s {
		//t.SchoolType = models.SchoolTypeHighSchool
		//err = models.GetSchoolDao().Save(*t)
		fmt.Println(t)
	}
	rsp.UserName = user.Name
	rsp.UserAge = user.Age
	return nil
}
