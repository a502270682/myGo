package log

import (
	"context"
	"io"
	"time"

	"github.com/sirupsen/logrus"
)

// XdLogger的接口
type MyLogger interface {
	SetOutput(out io.Writer)
	SetFormatter(formatter Formatter)
	SetReportCaller(include bool)
	SetLevel(level logrus.Level)
	AddHook(hook logrus.Hook)
	WithField(key string, value interface{}) LoggerEntry
	WithFields(fields Fields) LoggerEntry
	WithError(err error) LoggerEntry
	WithTime(t time.Time) LoggerEntry
	WithObject(obj interface{}) LoggerEntry
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

type CtxLogger struct {
	n *logrus.Logger // normal log
}

func (cl *CtxLogger) newLogEntry() LoggerEntry {
	return &LogEntry{logrus.NewEntry(cl.n), cl.n}
}

func (cl *CtxLogger) SetOutput(out io.Writer) {
	cl.n.SetOutput(out)
}

func (cl *CtxLogger) SetFormatter(formatter Formatter) {
	cl.n.SetFormatter(formatter)
}

func (cl *CtxLogger) SetReportCaller(include bool) {
	cl.n.SetReportCaller(include)
}

func (cl *CtxLogger) SetLevel(level logrus.Level) {
	cl.n.SetLevel(level)
}

func (cl *CtxLogger) AddHook(hook logrus.Hook) {
	cl.n.AddHook(hook)
}

func (cl *CtxLogger) WithField(key string, value interface{}) LoggerEntry {
	// 借用logrus.Logger本身Entry的管理机制来创建Entry,下同
	return &LogEntry{cl.n.WithField(key, value), cl.n}
}

func (cl *CtxLogger) WithFields(fields Fields) LoggerEntry {
	return &LogEntry{cl.n.WithFields(fields), cl.n}
}

func (cl *CtxLogger) WithError(err error) LoggerEntry {
	return &LogEntry{cl.n.WithError(err), cl.n}
}

func (cl *CtxLogger) WithTime(t time.Time) LoggerEntry {
	return &LogEntry{cl.n.WithTime(t), cl.n}
}

func (cl *CtxLogger) WithObject(obj interface{}) LoggerEntry {
	fields := parseFieldsFromObj(obj)
	return &LogEntry{cl.n.WithFields(fields), cl.n}
}

func (cl *CtxLogger) Tracef(ctx context.Context, format string, args ...interface{}) {
	cl.newLogEntry().Tracef(ctx, format, args...)
}

func (cl *CtxLogger) Debugf(ctx context.Context, format string, args ...interface{}) {
	cl.newLogEntry().Debugf(ctx, format, args...)
}

func (cl *CtxLogger) Infof(ctx context.Context, format string, args ...interface{}) {
	cl.newLogEntry().Infof(ctx, format, args...)
}

func (cl *CtxLogger) Printf(ctx context.Context, format string, args ...interface{}) {
	cl.newLogEntry().Printf(ctx, format, args...)
}

func (cl *CtxLogger) Warnf(ctx context.Context, format string, args ...interface{}) {
	cl.newLogEntry().Warnf(ctx, format, args...)
}

func (cl *CtxLogger) Warningf(ctx context.Context, format string, args ...interface{}) {
	cl.newLogEntry().Warningf(ctx, format, args...)
}

func (cl *CtxLogger) Errorf(ctx context.Context, format string, args ...interface{}) {
	cl.newLogEntry().Errorf(ctx, format, args...)
}

func (cl *CtxLogger) Fatalf(ctx context.Context, format string, args ...interface{}) {
	cl.newLogEntry().Fatalf(ctx, format, args...)
}

func (cl *CtxLogger) Panicf(ctx context.Context, format string, args ...interface{}) {
	cl.newLogEntry().Panicf(ctx, format, args...)
}

func (cl *CtxLogger) Log(ctx context.Context, level Level, args ...interface{}) {
	cl.newLogEntry().Log(ctx, level, args...)
}

func (cl *CtxLogger) Trace(ctx context.Context, args ...interface{}) {
	cl.newLogEntry().Trace(ctx, args...)
}

func (cl *CtxLogger) Debug(ctx context.Context, args ...interface{}) {
	cl.newLogEntry().Debug(ctx, args...)
}

func (cl *CtxLogger) Info(ctx context.Context, args ...interface{}) {
	cl.newLogEntry().Info(ctx, args...)
}

func (cl *CtxLogger) Print(ctx context.Context, args ...interface{}) {
	cl.newLogEntry().Print(ctx, args...)
}

func (cl *CtxLogger) Warn(ctx context.Context, args ...interface{}) {
	cl.newLogEntry().Warn(ctx, args...)
}

func (cl *CtxLogger) Warning(ctx context.Context, args ...interface{}) {
	cl.newLogEntry().Warning(ctx, args...)
}

func (cl *CtxLogger) Error(ctx context.Context, args ...interface{}) {
	cl.newLogEntry().Error(ctx, args...)
}

func (cl *CtxLogger) Fatal(ctx context.Context, args ...interface{}) {
	cl.newLogEntry().Fatal(ctx, args...)
}

func (cl *CtxLogger) Panic(ctx context.Context, args ...interface{}) {
	cl.newLogEntry().Panic(ctx, args...)
}

func (cl *CtxLogger) Logln(ctx context.Context, level Level, args ...interface{}) {
	cl.newLogEntry().Logln(ctx, level, args...)
}

func (cl *CtxLogger) Traceln(ctx context.Context, args ...interface{}) {
	cl.newLogEntry().Traceln(ctx, args...)
}

func (cl *CtxLogger) Debugln(ctx context.Context, args ...interface{}) {
	cl.newLogEntry().Debugln(ctx, args...)
}

func (cl *CtxLogger) Infoln(ctx context.Context, args ...interface{}) {
	cl.newLogEntry().Infoln(ctx, args...)
}

func (cl *CtxLogger) Println(ctx context.Context, args ...interface{}) {
	cl.newLogEntry().Println(ctx, args...)
}

func (cl *CtxLogger) Warnln(ctx context.Context, args ...interface{}) {
	cl.newLogEntry().Warnln(ctx, args...)
}

func (cl *CtxLogger) Warningln(ctx context.Context, args ...interface{}) {
	cl.newLogEntry().Warningln(ctx, args...)
}

func (cl *CtxLogger) Errorln(ctx context.Context, args ...interface{}) {
	cl.newLogEntry().Errorln(ctx, args...)
}

func (cl *CtxLogger) Fatalln(ctx context.Context, args ...interface{}) {
	cl.newLogEntry().Fatalln(ctx, args...)
}

func (cl *CtxLogger) Panicln(ctx context.Context, args ...interface{}) {
	cl.newLogEntry().Panicln(ctx, args...)
}
