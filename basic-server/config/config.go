package config

import (
	"fmt"

	"github.com/spf13/viper"
)

const defaultPort = "1234"
const defaultEnvironment = "development"
const environmentPrefix = "BASIC_SERVER"

// Init initialises the configuration
// 	when in development, load the ./config.yaml
// 	when in production, load from the environment
// 	environment variables should be prefixed with BASIC_SERVER_*
func Init() {
	viper.SetDefault("port", defaultPort)
	viper.SetDefault("env", defaultEnvironment)
	viper.SetEnvPrefix(environmentPrefix)
	viper.BindEnv("env")
	if viper.GetString("env") == "development" {
		viper.SetConfigName("config")
		viper.AddConfigPath(".")
		err := viper.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("%s", err))
		}
	} else {
		viper.BindEnv("port")
	}
}
