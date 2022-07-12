package types

type Logger interface {
	Infof(format string, v ...interface{})
	Errorf(format string, v ...interface{})
	Warnf(format string, v ...interface{})
	Fatalf(format string, v ...interface{})
	Debugf(format string, v ...interface{})
	Info(args ...interface{})
	Error(args ...interface{})
	Warn(args ...interface{})
	Fatal(args ...interface{})
	Debug(args ...interface{})
}
