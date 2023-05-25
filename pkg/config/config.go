package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port           string `mapstructure:"PORT"`
	AuthServiceUrl string `mapstructure:"AUTH_SERVICE_URL"`
}

func LoadConfig() (config *Config, err error) {

	viper.AddConfigPath("./pkg/config/envs")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	if err = viper.ReadInConfig(); err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
