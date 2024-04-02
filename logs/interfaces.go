package logs

type LogHandler func(string) error

type Logger interface {
	Debug(string, ...interface{})
	Error(string, ...interface{})
	ErrorStack(err error)
	Info(string, ...interface{})
	Fatal(string, ...interface{})
	Warn(string, ...interface{})
	Log(string, Level, ...interface{})
}

type Log interface {
	Debug(string) string
	Error(string) string
	Info(string) string
	Fatal(string) string
	Warn(string) string
}

type LogHook interface {
	OnLog(string, Level)
}

type LoggerCreator interface {
	Create(options *LogOptions) (Log, error)
}
