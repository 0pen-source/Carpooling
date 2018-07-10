package main

import (
	"os"

	"github.com/0pen-source/Carpooling/models"
	"github.com/0pen-source/Carpooling/utils"

	"github.com/spf13/viper"
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
}
