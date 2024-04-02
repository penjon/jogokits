package logs

import (
	"fmt"
	"github.com/penjon/jogokits/utils"
)

type LogFactory struct {
	log     *proxy
	options *LogOptions
	creator LoggerCreator
}

func (f *LogFactory) SetLogCreator(creator LoggerCreator, options *LogOptions) error {
	f.creator = creator
	log, err := creator.Create(options)
	if err != nil {
		return err
	}
	f.log = &proxy{
		logger:  log,
		options: options,
	}
	return nil
}

func (f *LogFactory) SetDefaultOptions(options *LogOptions) error {
	return f.SetLogCreator(&zapLogCreator{}, options)
}

func (f *LogFactory) getLogger() Logger {
	if f.log == nil {
		creator := &zapLogCreator{}
		if e := f.SetLogCreator(creator, &LogOptions{
			FileName:     utils.GetLogNameByAppName(),
			FileSize:     10,
			MaxFileCount: 30,
			Debug:        false,
			Compress:     true,
			Json:         false,
		}); e != nil {
			fmt.Printf("log creator set error.error[%s]", e.Error())
		}
	}
	return f.log
}

var factory *LogFactory

func GetFactory() *LogFactory {
	if nil == factory {
		factory = &LogFactory{}
	}
	return factory
}

func Get() Logger {
	return GetFactory().getLogger()
}
