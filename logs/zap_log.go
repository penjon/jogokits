package logs

import (
	"fmt"
	"github.com/penjon/jogokits/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

type zapLogCreator struct {
}

type zapLogger struct {
	app     *zap.Logger
	options *LogOptions
}

func (l *zapLogger) Debug(msg string) string {
	l.app.Debug(msg)
	return msg
}

func (l *zapLogger) Error(msg string) string {
	l.app.Error(msg)
	return msg
}

func (l *zapLogger) Info(msg string) string {
	l.app.Info(msg)
	return msg
}

func (l *zapLogger) Fatal(msg string) string {
	l.app.Fatal(msg)
	return msg
}

func (l *zapLogger) Warn(msg string) string {
	l.app.Warn(msg)
	return msg
}

func (z *zapLogCreator) Create(options *LogOptions) (Log, error) {
	log := &zapLogger{
		options: options,
	}
	log.app = NewLogger(options)

	return log, nil
}

/**
 * 获取日志
 * filePath 日志文件路径
 * level 日志级别
 * maxSize 每个日志文件保存的最大尺寸 单位：M
 * maxBackups 日志文件最多保存多少个备份
 * maxAge 文件最多保存多少天
 * compress 是否压缩
 * serviceName 服务名
 */
func NewLogger(ops *LogOptions) *zap.Logger {
	core := newCoreByOps(ops)
	return zap.New(core, zap.AddCaller(), zap.Development(), zap.AddCallerSkip(2))
}

func newCoreByOps(options *LogOptions) zapcore.Core {
	lv := zap.InfoLevel
	if options.Debug {
		lv = zap.DebugLevel
	}
	fileName := options.FileName
	if len(fileName) == 0 {
		fileName = utils.GetLogNameByAppName()
	}
	//如果有设置文件路径则处理路径
	if len(options.FilePath) >= 0 {
		fileName = fmt.Sprintf("%s/%s", options.FilePath, fileName)
	}
	//日志文件路径配置2
	hook := lumberjack.Logger{
		Filename:   fileName,             // 日志文件路径
		MaxSize:    options.FileSize,     // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: options.MaxFileCount, // 日志文件最多保存多少个备份
		MaxAge:     60,                   // 文件最多保存多少天
		Compress:   options.Compress,     // 是否压缩
	}
	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(lv)

	//公用编码器
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "linenum",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.ShortCallerEncoder,     // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}

	if options.Color {
		encoderConfig.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	}

	var encoder zapcore.Encoder
	if options.Json {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	return zapcore.NewCore(encoder,
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout),
			zapcore.AddSync(&hook)), // 打印到控制台和文件
		atomicLevel, // 日志级别
	)
}
