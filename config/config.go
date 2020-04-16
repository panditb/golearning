package config

import (
	"log"

	"github.com/spf13/viper"
)

func GetConfig() Configuration {

	//Check yam file in current directory
	viper.SetConfigName("config/config")
	viper.AddConfigPath(".")

	var configuration Configuration
	// Find and read config file
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	log.Printf("database uri is %s", configuration.Port)
	log.Printf("port for this application is %d", configuration.ConsumedUrl)

	return configuration
}
