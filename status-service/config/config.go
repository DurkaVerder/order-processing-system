package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type Endpoint struct {
	Base      string            `mapstructure:"base"`
	Endpoints map[string]string `mapstructure:"endpoints"`
}

type Service struct {
	Route  Endpoint `mapstructure:"route"`
	Server struct {
		Port string `mapstructure:"port"`
	} `mapstructure:"server"`
}

type Config struct {
	Status Service
}

func LoadConfig() Config {
	viper.SetConfigFile(os.Getenv("CONFIG_PATH"))

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Unable to read config: %v", err)
	}

	var status Service
	if err := viper.UnmarshalKey("Status", &status); err != nil {
		log.Fatalf("Unable to unmarshal Authentication config: %v", err)
	}

	log.Println("Config loaded successfully")

	return Config{
		Status: status,
	}
}
