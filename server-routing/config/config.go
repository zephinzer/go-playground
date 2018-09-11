package config

import (
	"server-routing/logger"

	"github.com/spf13/viper"
)

func Init() {
	viper.SetDefault("port", "1234")
	viper.SetDefault("env", "development")
	viper.BindEnv("env")
	if viper.GetString("env") == "development" {
		viper.SetConfigName("config")
		viper.AddConfigPath(".")
		err := viper.ReadInConfig()
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		viper.BindEnv("port")
	}
}
