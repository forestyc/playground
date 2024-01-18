package zap

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

type Config struct {
	Level         string `mapstructure:"level"`
	Format        string `mapstructure:"format"`
	Director      string `mapstructure:"director"`
	LinkName      string `mapstructure:"link-name"`
	ShowLine      bool   `mapstructure:"show-line"`
	EncodeLevel   string `mapstructure:"encode-level"`
	StacktraceKey string `mapstructure:"stacktrace-key"`
	LogInConsole  bool   `mapstructure:"log-in-console"`
}

type Zap struct {
	*zap.Logger
}

func NewZap(z Config) *Zap {
	mkdir(z.Director)
	var level zapcore.Level
	switch z.Level {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	case "dpanic":
		level = zap.DPanicLevel
	case "panic":
		level = zap.PanicLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}
	var logger Zap
	if level == zap.DebugLevel || level == zap.ErrorLevel {
		logger.Logger = zap.New(newCore(z, level), zap.AddStacktrace(level))
	} else {
		logger.Logger = zap.New(newCore(z, level))
	}
	if z.ShowLine {
		logger.Logger = logger.WithOptions(zap.AddCaller())
	}
	return &logger
}

// new zapcore.EncoderConfig
func newEncoderConfig(z Config) (config zapcore.EncoderConfig) {
	config = zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  z.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     newTimeEncoder(),
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
	switch {
	case z.EncodeLevel == "LowercaseLevelEncoder": //小写编码器（默认）
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	case z.EncodeLevel == "LowercaseColorLevelEncoder": //小写编码器带颜色
		config.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	case z.EncodeLevel == "CapitalLevelEncoder": //大写编码器
		config.EncodeLevel = zapcore.CapitalLevelEncoder
	case z.EncodeLevel == "CapitalColorLevelEncoder": //大写编码器带颜色
		config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	default:
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	}
	return config
}

// new zapcore.Core
func newCore(z Config, level zapcore.Level) (core zapcore.Core) {
	writer, err := GetWriteSyncer(z) //使用file-rotatelogs进行日志分割
	if err != nil {
		fmt.Printf("Get Write Syncer Failed err:%v", err)
		return
	}
	var encoder zapcore.Encoder
	if z.Format == "json" {
		encoder = zapcore.NewJSONEncoder(newEncoderConfig(z))
	} else {
		encoder = zapcore.NewConsoleEncoder(newEncoderConfig(z))
	}
	return zapcore.NewCore(encoder, writer, level)
}

// new zapcore.TimeEncoder
func newTimeEncoder() zapcore.TimeEncoder {
	return func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006/01/02 - 15:04:05.000"))
	}
}

// create log path
func mkdir(path string) {
	os.Mkdir(path, os.ModePerm)
}
