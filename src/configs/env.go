package configs

import (
	"log"

	"github.com/spf13/viper"
)

var EnvConfigs *envConfig

func InitEnvConfig() {
	EnvConfigs = loadEnvVariables()
}

type envConfig struct {
	LocalServerPort string `mapstructure:"LOCAL_SERVER_PORT"`
	DbPassword      string `mapstructure:"DB_PASSWORD"`
	DbUsername      string `mapstructure:"DB_USERNAME"`
}

func loadEnvVariables() (config *envConfig) {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error getting env file", err)
	}
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal("Error unmarshalling config", err)
	}
	return
}
