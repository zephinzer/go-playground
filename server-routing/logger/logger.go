package logger

import (
	zap "go.uber.org/zap"
)

var logger, _ = zap.NewProduction()
var loggerInstace = logger.Sugar()

var Debug = loggerInstace.Debug
var Info = loggerInstace.Info
var Infof = loggerInstace.Infof
var Infow = loggerInstace.Infow
var Warn = loggerInstace.Warn
var Warnf = loggerInstace.Warnf
var Warnw = loggerInstace.Warnw
var Error = loggerInstace.Error
var Errorf = loggerInstace.Errorf
var Errorw = loggerInstace.Errorw
var DPanic = loggerInstace.DPanic
var Panic = loggerInstace.Panic
var Panicf = loggerInstace.Panicf
var Panicw = loggerInstace.Panicw
var Fatal = loggerInstace.Fatal
var Fatalf = loggerInstace.Fatalf
var Fatalw = loggerInstace.Fatalw
