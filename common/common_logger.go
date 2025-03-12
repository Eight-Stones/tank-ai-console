package common

type LoggerI interface {
	Debug(args ...any)
	Info(args ...any)
	Warn(args ...any)
	Error(args ...any)
	Debugf(template string, args ...any)
	Infof(template string, args ...any)
	Warnf(template string, args ...any)
	Errorf(template string, args ...any)
	Print(args ...any)
	Printf(format string, args ...any)
	Name(name string) LoggerI
	Fields(fields ...any) LoggerI
}
