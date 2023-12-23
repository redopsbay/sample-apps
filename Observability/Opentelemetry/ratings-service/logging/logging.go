package logging

import (
	"go.uber.org/zap"
)

func Info(logs interface{}) {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	logging := logger.Sugar()
	logging.Info(logs)
}

func Warn(logs interface{}) {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	logging := logger.Sugar()
	logging.Warn(logs)
}

func Error(logs interface{}) {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	logging := logger.Sugar()
	logging.Error(logs)
}

func Fatal(logs interface{}) {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	logging := logger.Sugar()
	logging.Fatal(logs)
}
