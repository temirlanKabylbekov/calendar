package logger

import "go.uber.org/zap/zapcore"

type LogMessage struct {
	LevelKey   string
	MessageKey string
	CallerKey  string
	TimeKey    string
}

func (m LogMessage) getLevelFormatter() zapcore.LevelEncoder {
	return zapcore.CapitalLevelEncoder
}

func (m LogMessage) getCallerFormatter() zapcore.CallerEncoder {
	return zapcore.ShortCallerEncoder
}

func (m LogMessage) getTimeFormatter() zapcore.TimeEncoder {
	return zapcore.ISO8601TimeEncoder
}
