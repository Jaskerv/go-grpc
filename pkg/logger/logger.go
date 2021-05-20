package logger

import (
	"log"
	"os"

	"go.uber.org/zap"
)

var Logger zap.SugaredLogger

func init() {
	var initFunc func(options ...zap.Option) (*zap.Logger, error)

	if ENV := os.Getenv("ENV"); ENV == "production" {
		initFunc = zap.NewProduction
	} else {
		initFunc = zap.NewDevelopment
	}

	logger, err := initFunc()
	if err != nil {
		log.Fatalln("Fatal: failed to initialise logger")
	}

	Logger := logger.Sugar()

	defer Logger.Sync()
}
