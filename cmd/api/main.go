package main

import (
	"github.com/rinatkh/artstudio_back/config"
	"github.com/rinatkh/artstudio_back/internal/httpServer"
	"github.com/rinatkh/artstudio_back/pkg/httpErrorHandler"
	"github.com/sirupsen/logrus"
	"log"
	"time"
)

func main() {

	v, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Cannot load config: ", err.Error())
	}
	cfg, err := config.ParseConfig(v)
	if err != nil {
		log.Fatal("Config parse error", err.Error())
	}

	logger := logrus.New()
	formatter := logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
	}
	switch cfg.Logger.Level {
	case "warning":
		logger.SetLevel(logrus.WarnLevel)
	case "notice":
		logger.SetLevel(logrus.InfoLevel)
	case "debug":
		logger.SetLevel(logrus.DebugLevel)
		formatter.PrettyPrint = true
	default:
		logger.SetLevel(logrus.InfoLevel)
	}
	logger.SetFormatter(&formatter)
	logger.Infof("logger level: %s", logger.Level.String())

	logger.Println("Config loaded")

	errorHandler := httpErrorHandler.NewErrorHandler(cfg)
	s := httpServer.NewServer(cfg, errorHandler, logrus.NewEntry(logger))
	if err = s.Run(); err != nil {
		logger.Fatalln(err)
	}
}
