package log

type Logger interface {
	Debugf(format string, v ...interface{})
	Infof(format string, v ...interface{})
	Warnf(format string, v ...interface{})
	Errorf(format string, v ...interface{})
	Fatalf(format string, v ...interface{})

	Infow(msg string, fields ...interface{})
	Warnw(msg string, fields ...interface{})
	Errorw(msg string, fields ...interface{})
	Fatalw(msg string, fields ...interface{})
	Debugw(msg string, fields ...interface{})
}
