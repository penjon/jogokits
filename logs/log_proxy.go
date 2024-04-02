package logs

import (
	"fmt"
)

type proxy struct {
	options *LogOptions
	hooks   []LogHook
	logger  Log
}

func (l *proxy) AddHook(hook LogHook) {
	if l.hooks == nil {
		l.hooks = make([]LogHook, 0)
	}
	l.hooks = append(l.hooks, hook)
}

func (l *proxy) Debug(msg string, args ...interface{}) {
	m := msg
	if nil != args && len(args) > 0 {
		m = fmt.Sprintf(msg, args...)
	}
	l.syncHook(m, DEBUG)
	l.logger.Debug(m)
}

func (l *proxy) Error(msg string, args ...interface{}) {
	m := msg
	if nil != args && len(args) > 0 {
		m = fmt.Sprintf(msg, args...)
	}
	l.syncHook(m, ERROR)
	l.logger.Error(m)
}
func (l *proxy) ErrorStack(err error) {
	m := fmt.Sprintf("%+v\n", err)
	l.syncHook(m, ERROR)
	l.logger.Error(m)
}
func (l *proxy) Fatal(msg string, args ...interface{}) {
	m := msg
	if nil != args && len(args) > 0 {
		m = fmt.Sprintf(msg, args...)
	}
	l.syncHook(m, FATAL)
	l.logger.Fatal(m)
}

func (l *proxy) Warn(msg string, args ...interface{}) {
	m := msg
	if nil != args && len(args) > 0 {
		m = fmt.Sprintf(msg, args...)
	}
	l.syncHook(m, WARN)
	l.logger.Warn(m)
}

func (l *proxy) Log(msg string, lv Level, args ...interface{}) {
	m := msg
	if nil != args && len(args) > 0 {
		m = fmt.Sprintf(msg, args...)
	}
	switch lv {
	case DEBUG:
		l.logger.Debug(m)
		break
	case INFO:
		l.logger.Info(m)
		break
	case ERROR:
		l.logger.Error(m)
		break
	case FATAL:
		l.logger.Fatal(m)
		break
	case WARN:
		l.logger.Warn(m)
	}
}

func (l *proxy) Info(msg string, args ...interface{}) {
	m := msg
	if nil != args && len(args) > 0 {
		m = fmt.Sprintf(msg, args...)
	}
	l.syncHook(m, INFO)
	l.logger.Info(m)
}

func (l *proxy) syncHook(msg string, lv Level) {
	if nil != l.hooks {
		for _, hook := range l.hooks {
			hook.OnLog(msg, lv)
		}
	}
}
