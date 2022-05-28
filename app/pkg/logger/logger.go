package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger interface {
	Infow(msg string, args ...interface{})
	Warnw(msg string, args ...interface{})
	Errorw(msg string, args ...interface{})

	Sync() error
}

type Config struct {
	WS               zapcore.WriteSyncer
	TimeEncoder      zapcore.TimeEncoder
	IsConsoleEncoder bool
}

func New(cfg *Config) Logger {
	encoderCfg := zap.NewProductionEncoderConfig()

	if cfg.TimeEncoder != nil {
		encoderCfg.EncodeTime = cfg.TimeEncoder
	} else {
		encoderCfg.EncodeTime = zapcore.RFC3339TimeEncoder
	}

	var ws zapcore.WriteSyncer
	if cfg.WS == nil {
		ws = zapcore.Lock(os.Stdout)
	} else {
		ws = zapcore.Lock(cfg.WS)
	}

	var encoder zapcore.Encoder
	if cfg.IsConsoleEncoder {
		encoder = zapcore.NewConsoleEncoder(encoderCfg)
	} else {
		encoder = zapcore.NewJSONEncoder(encoderCfg)
	}

	return zap.New(zapcore.NewCore(encoder, ws, zap.NewAtomicLevel()), zap.AddCallerSkip(1)).Sugar()
}
