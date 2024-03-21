package main

import (
	"errors"
	"gituhb.com/ropehapi/logger-go/pkg/logger"
	"go.uber.org/zap"
)

func main() {
	logger.Info("Your message here",
		zap.String("key", "value"),
	)

	logger.Error("Your error here",
		errors.New("Your error here"),
		zap.String("key", "value"),
	)
}
