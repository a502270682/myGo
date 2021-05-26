package log

import (
	"context"
	"reflect"
	"time"

	"github.com/sirupsen/logrus"
)

// 实现接口 MyLoggerEntry
type MyGoLogEntry struct {
	*logrus.Entry
	nl *logrus.Logger
	sl *logrus.Logger
}

// log 以及 entry的公用接口，方法一致，业务使用体验一致
type MyLoggerEntry interface {
	WithField(key string, value interface{}) MyLoggerEntry
	WithFields(fields Fields) MyLoggerEntry
	WithError(err error) MyLoggerEntry
	WithTime(t time.Time) MyLoggerEntry
	WithObject(obj interface{}) MyLoggerEntry
	Tracef(ctx context.Context, format string, args ...interface{})
	Debugf(ctx context.Context, format string, args ...interface{})
	Infof(ctx context.Context, format string, args ...interface{})
	Printf(ctx context.Context, format string, args ...interface{})
	Warnf(ctx context.Context, format string, args ...interface{})
	Warningf(ctx context.Context, format string, args ...interface{})
	Errorf(ctx context.Context, format string, args ...interface{})
	Fatalf(ctx context.Context, format string, args ...interface{})
	Panicf(ctx context.Context, format string, args ...interface{})
	Log(ctx context.Context, level Level, args ...interface{})
	Trace(ctx context.Context, args ...interface{})
	Debug(ctx context.Context, args ...interface{})
	Info(ctx context.Context, args ...interface{})
	Print(ctx context.Context, args ...interface{})
	Warn(ctx context.Context, args ...interface{})
	Warning(ctx context.Context, args ...interface{})
	Error(ctx context.Context, args ...interface{})
	Fatal(ctx context.Context, args ...interface{})
	Panic(ctx context.Context, args ...interface{})
	Logln(ctx context.Context, level Level, args ...interface{})
	Traceln(ctx context.Context, args ...interface{})
	Debugln(ctx context.Context, args ...interface{})
	Infoln(ctx context.Context, args ...interface{})
	Println(ctx context.Context, args ...interface{})
	Warnln(ctx context.Context, args ...interface{})
	Warningln(ctx context.Context, args ...interface{})
	Errorln(ctx context.Context, args ...interface{})
	Fatalln(ctx context.Context, args ...interface{})
	Panicln(ctx context.Context, args ...interface{})
}

func (en MyGoLogEntry) WithField(key string, value interface{}) MyLoggerEntry {
	return &MyGoLogEntry{en.Entry.WithField(key, value), en.nl, en.sl}
}

func (en MyGoLogEntry) WithFields(fields Fields) MyLoggerEntry {
	return &MyGoLogEntry{en.Entry.WithFields(fields), en.nl, en.sl}
}

func (en MyGoLogEntry) WithError(err error) MyLoggerEntry {
	return &MyGoLogEntry{en.Entry.WithError(err), en.nl, en.sl}
}

func (en MyGoLogEntry) WithTime(t time.Time) MyLoggerEntry {
	return &MyGoLogEntry{en.Entry.WithTime(t), en.nl, en.sl}
}

func (en MyGoLogEntry) WithObject(obj interface{}) MyLoggerEntry {
	fields := parseFieldsFromObj(obj)
	return &MyGoLogEntry{en.Entry.WithFields(fields), en.nl, en.sl}
}

func (en MyGoLogEntry) Tracef(ctx context.Context, format string, args ...interface{}) {
	en.Entry.WithContext(ctx).Logf(TraceLevel, format, args...)
}

func (en MyGoLogEntry) Debugf(ctx context.Context, format string, args ...interface{}) {
	en.Entry.WithContext(ctx).Logf(DebugLevel, format, args...)
}

func (en MyGoLogEntry) Infof(ctx context.Context, format string, args ...interface{}) {
	en.Entry.WithContext(ctx).Logf(InfoLevel, format, args...)
}

func (en MyGoLogEntry) Printf(ctx context.Context, format string, args ...interface{}) {

	en.Entry.WithContext(ctx).Printf(format, args...)
}

func (en MyGoLogEntry) Warnf(ctx context.Context, format string, args ...interface{}) {

	en.Entry.WithContext(ctx).Logf(WarnLevel, format, args...)
}

func (en MyGoLogEntry) Warningf(ctx context.Context, format string, args ...interface{}) {

	en.Entry.WithContext(ctx).Warnf(format, args...)
}

func (en MyGoLogEntry) Errorf(ctx context.Context, format string, args ...interface{}) {

	en.Entry.WithContext(ctx).Logf(ErrorLevel, format, args...)
}

func (en MyGoLogEntry) Fatalf(ctx context.Context, format string, args ...interface{}) {

	en.Entry.WithContext(ctx).Fatalf(format, args...)
}

func (en MyGoLogEntry) Panicf(ctx context.Context, format string, args ...interface{}) {

	en.Entry.WithContext(ctx).Logf(PanicLevel, format, args...)
}

func (en MyGoLogEntry) Log(ctx context.Context, level Level, args ...interface{}) {

	en.Entry.WithContext(ctx).Log(level, args...)
}

func (en MyGoLogEntry) Trace(ctx context.Context, args ...interface{}) {

	en.Entry.WithContext(ctx).Log(TraceLevel, args...)
}

func (en MyGoLogEntry) Debug(ctx context.Context, args ...interface{}) {

	en.Entry.WithContext(ctx).Log(DebugLevel, args...)
}

func (en MyGoLogEntry) Info(ctx context.Context, args ...interface{}) {

	en.Entry.WithContext(ctx).Log(InfoLevel, args...)
}

func (en MyGoLogEntry) Print(ctx context.Context, args ...interface{}) {

	en.Entry.WithContext(ctx).Print(args...)
}

func (en MyGoLogEntry) Warn(ctx context.Context, args ...interface{}) {

	en.Entry.WithContext(ctx).Log(WarnLevel, args...)
}

func (en MyGoLogEntry) Warning(ctx context.Context, args ...interface{}) {

	en.Entry.WithContext(ctx).Warn(args...)
}

func (en MyGoLogEntry) Error(ctx context.Context, args ...interface{}) {

	en.Entry.WithContext(ctx).Log(ErrorLevel, args...)
}

func (en MyGoLogEntry) Fatal(ctx context.Context, args ...interface{}) {

	en.Entry.WithContext(ctx).Fatal(args...)
}

func (en MyGoLogEntry) Panic(ctx context.Context, args ...interface{}) {

	en.Entry.WithContext(ctx).Panic(args...)
}

func (en MyGoLogEntry) Logln(ctx context.Context, level Level, args ...interface{}) {

	en.Entry.WithContext(ctx).Logln(level, args...)
}

func (en MyGoLogEntry) Traceln(ctx context.Context, args ...interface{}) {

	en.Entry.WithContext(ctx).Logln(TraceLevel, args...)
}

func (en MyGoLogEntry) Debugln(ctx context.Context, args ...interface{}) {

	en.Entry.WithContext(ctx).Logln(DebugLevel, args...)
}

func (en MyGoLogEntry) Infoln(ctx context.Context, args ...interface{}) {

	en.Entry.WithContext(ctx).Logln(InfoLevel, args...)
}

func (en MyGoLogEntry) Println(ctx context.Context, args ...interface{}) {

	en.Entry.WithContext(ctx).Println(args...)
}

func (en MyGoLogEntry) Warnln(ctx context.Context, args ...interface{}) {

	en.Entry.WithContext(ctx).Logln(WarnLevel, args...)
}

func (en MyGoLogEntry) Warningln(ctx context.Context, args ...interface{}) {

	en.Entry.WithContext(ctx).Logln(WarnLevel, args...)
}

func (en MyGoLogEntry) Errorln(ctx context.Context, args ...interface{}) {

	en.Entry.WithContext(ctx).Logln(ErrorLevel, args...)
}

func (en MyGoLogEntry) Fatalln(ctx context.Context, args ...interface{}) {

	en.Entry.WithContext(ctx).Fatalln(args...)
}

func (en MyGoLogEntry) Panicln(ctx context.Context, args ...interface{}) {

	en.Entry.WithContext(ctx).Logln(PanicLevel, args...)
}

func isZero(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.String:
		return len(v.String()) == 0
	case reflect.Ptr, reflect.Interface:
		return v.IsNil()
	case reflect.Slice:
		return v.Len() == 0
	case reflect.Map:
		return v.Len() == 0
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Struct: // 不去确认
		return false
	}
	return false
}

func parseFieldsFromObj(o interface{}) logrus.Fields {
	logFields := logrus.Fields{}

	val := reflect.ValueOf(o)
	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return logFields
		}
		val = val.Elem()
	}
	for i := 0; i < val.NumField(); i++ {
		fValue := val.Field(i)
		fType := val.Type().Field(i)
		if !isZero(fValue) && fValue.IsValid() && fType.PkgPath == "" { // exported fields
			if fValue.Kind() == reflect.Struct ||
				(fValue.Kind() == reflect.Ptr &&
					fValue.Elem().Kind() == reflect.Struct) {
				fields := parseFieldsFromObj(fValue.Interface())
				if fType.Anonymous {
					for k, v := range fields {
						logFields[k] = v
					}
				} else {
					logFields[fType.Name] = fields
				}
			} else {
				logFields[fType.Name] = fValue.Interface()
			}
		}
	}
	return logFields
}
