package main

import (
	"log"

	"github.com/temirlanKabylbekov/calendar/cmd"
	"github.com/temirlanKabylbekov/calendar/config"
	"github.com/temirlanKabylbekov/calendar/logger"
)

const (
	configFileName = "config"
	configFilePath = "."
)

var cfg *config.Config

func init() {
	loader := config.ConfigLoader{FileName: configFileName, FilePath: configFilePath}
	_cfg, err := loader.Load()
	if err != nil {
		log.Fatal(err)
	}
	cfg = _cfg
}

func getLogger() (*logger.Logger, error) {
	msgType := logger.LogMessage{
		LevelKey:   cfg.Logger.Message.LevelName,
		MessageKey: cfg.Logger.Message.MessageName,
		CallerKey:  cfg.Logger.Message.CallerName,
		TimeKey:    cfg.Logger.Message.TimeName,
	}
	loggerConfig := logger.LoggerConfig{
		Level:            cfg.Logger.Level,
		Encoding:         cfg.Logger.Encoding,
		MessageType:      msgType,
		Development:      cfg.Development,
		OutputPaths:      []string{cfg.Logger.FilePath},
		ErrorOutputPaths: []string{cfg.Logger.FilePath},
	}
	logger, err := logger.NewLogger(loggerConfig)
	return logger, err
}

func main() {
	Log, err := getLogger()
	if err != nil {
		log.Fatal(err)
	}

	server := cmd.Server{Address: cfg.Server.Address(), Log: Log}
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
