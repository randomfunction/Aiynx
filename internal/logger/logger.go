package logger

import (
	"log"
	"sync"

	"go.uber.org/zap"
)

var (
	Log  *zap.Logger
	once sync.Once
)

func InitLogger() {
	once.Do(func() {
		var err error
		Log, err = zap.NewProduction()
		if err != nil {
			log.Fatalf("Failed to initialize logger: %v", err)
		}
	})
}

func Sync() {
	if Log != nil {
		_ = Log.Sync()
	}
}
