package main

import (
	config "github.com/spf13/viper"
	zap "go.uber.org/zap"
)

type testObjectToLog struct {
	SomeInteger int    `json:"some-integer"`
	SomeString  string `json:"some-string"`
}

var productionLogger *zap.Logger
var developmentLogger *zap.SugaredLogger

func main() {
	initConfig()
	initLogger()
	demonstrateDebug(developmentLogger)
	demonstrateInfo(productionLogger)
	demonstrateWarn(productionLogger)
	demonstrateError(productionLogger)
	demonstrateFatal(productionLogger)
	return
}

func initConfig() {
	config.SetDefault("env", "development")
	config.BindEnv("env")
}

func initLogger() {
	zapProduction, _ := zap.NewProduction()
	productionLogger = zapProduction
	zapDevelopment, _ := zap.NewDevelopment()
	developmentLogger = zapDevelopment.Sugar()
}

func demonstrateDebug(logger *zap.SugaredLogger) {
	logger.Debug("[dev] logger.Debug")
	logger.Debugw("[dev] logger.Debugw", "log", testObjectToLog{
		SomeInteger: 1,
	})
	logger.Debugf("[dev] logger.Debugf %s", "log")
}

func demonstrateInfo(logger *zap.Logger) {
	logger.Info("logger.Info")
	logger.Info(
		"logger.Info",
		zap.Any("testObjectToLog",
			testObjectToLog{
				SomeInteger: 1,
			},
		),
	)
}

func demonstrateWarn(logger *zap.Logger) {
	logger.Warn("logger.Warn")
	logger.Warn(
		"logger.Warn",
		zap.Any("testObjectToLog",
			testObjectToLog{
				SomeInteger: 1,
			},
		),
	)
}

func demonstrateError(logger *zap.Logger) {
	logger.Error("logger.Error")
	logger.Error(
		"logger.Error",
		zap.Any("testObjectToLog",
			testObjectToLog{
				SomeInteger: 1,
			},
		),
	)
}

func demonstrateFatal(logger *zap.Logger) {
	logger.Fatal("logger.Fatal")
	logger.Fatal(
		"logger.Fatal",
		zap.Any("testObjectToLog",
			testObjectToLog{
				SomeInteger: 1,
			},
		),
	)
}
