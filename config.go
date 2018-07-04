package main

import (
	"github.com/Carpooling/models"
	"github.com/Carpooling/utils"
	"github.com/spf13/viper"
	"os"
)

var config models.Configuration

func initializeConfiguration() {
	filename := os.Getenv("CARPOOLING_SERVER_CONF")
	if filename == "" {
		filename = "./conf.yml"
	}

	viper.SetConfigType("yaml")
	viper.SetConfigFile(filename)

	utils.Must(nil, viper.ReadInConfig())
	utils.Must(nil, viper.Unmarshal(&config))
	utils.Must(nil, config.Validate())
}
