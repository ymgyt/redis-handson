package core

import "go.uber.org/zap"

var zapLogger, err = zap.NewDevelopment()

var Logger = zapLogger.Sugar()
