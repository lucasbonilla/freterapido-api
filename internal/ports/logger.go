package ports

type Logger interface {
	Debug(value ...interface{})
	Info(value ...interface{})
	Warning(value ...interface{})
	Error(value ...interface{})
	Debugf(format string, v ...interface{})
	Infof(format string, v ...interface{})
	Warningf(format string, v ...interface{})
	Errorf(format string, v ...interface{})
}
