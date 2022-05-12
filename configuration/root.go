package configuration

import (
	"github.com/spf13/viper"
)

type Config struct {
	ServiceName string

	ServerPort string

	LogJson  bool
	LogLevel string
}

var config *Config

func GetConfig() *Config {
	if config == nil {
		config = &Config{
			ServiceName: viper.GetString("REST_API_SERVICE_NAME"),
			ServerPort:  viper.GetString("REST_API_SERVER_PORT"),

			LogJson:  viper.GetBool("REST_API_LOG_JSON"),
			LogLevel: viper.GetString("REST_API_LOG_LEVEL"),
		}
	}
	return config
}
