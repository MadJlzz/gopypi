package main

import (
	"go.uber.org/zap"
	"log"
)

func main() {
	l, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("can't initialize zap l: %v", err)
	}
	defer l.Sync()

	// SugaredLogger includes both printf-style APIs.
	logger := l.Sugar()
	logger.Info("Will list Google cloud stuff")

}
