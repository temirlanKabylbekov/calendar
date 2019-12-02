package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	inner *zap.SugaredLogger
}

func (l *Logger) Debug(args ...interface{}) {
	l.inner.Debug(args...)
}

func (l *Logger) Debugf(template string, args ...interface{}) {
	l.inner.Debugf(template, args...)
}

func (l *Logger) Debugw(msg string, keysAndValues ...interface{}) {
	l.inner.Debugw(msg, keysAndValues...)
}

func (l *Logger) Info(args ...interface{}) {
	l.inner.Info(args...)
}

func (l *Logger) Infof(template string, args ...interface{}) {
	l.inner.Infof(template, args...)
}

func (l *Logger) Infow(msg string, keysAndValues ...interface{}) {
	l.inner.Infow(msg, keysAndValues...)
}

func (l *Logger) Warn(args ...interface{}) {
	l.inner.Warn(args...)
}

func (l *Logger) Warnf(template string, args ...interface{}) {
	l.inner.Warnf(template, args...)
}

func (l *Logger) Warnw(msg string, keysAndValues ...interface{}) {
	l.inner.Warnw(msg, keysAndValues...)
}

func (l *Logger) Error(args ...interface{}) {
	l.inner.Error(args...)
}

func (l *Logger) Errorf(template string, args ...interface{}) {
	l.inner.Errorf(template, args...)
}

func (l *Logger) Errorw(msg string, keysAndValues ...interface{}) {
	l.inner.Errorw(msg, keysAndValues...)
}

func NewLogger(cfg LoggerConfig) (*Logger, error) {
	l, err := zap.Config{
		Level:            cfg.getLevel(),
		Encoding:         cfg.Encoding,
		Development:      cfg.Development,
		OutputPaths:      cfg.OutputPaths,
		ErrorOutputPaths: cfg.ErrorOutputPaths,
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:   cfg.MessageType.MessageKey,
			LevelKey:     cfg.MessageType.LevelKey,
			EncodeLevel:  cfg.MessageType.getLevelFormatter(),
			TimeKey:      cfg.MessageType.TimeKey,
			EncodeTime:   cfg.MessageType.getTimeFormatter(),
			CallerKey:    cfg.MessageType.CallerKey,
			EncodeCaller: cfg.MessageType.getCallerFormatter(),
		},
	}.Build()

	if err != nil {
		return &Logger{}, err
	}

	// добавили обертку, поэтому по умолчанию показывало бы местом вызова логирования функцию-обертки
	l = l.WithOptions(zap.AddCallerSkip(1))

	return &Logger{inner: l.Sugar()}, nil
}
