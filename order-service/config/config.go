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
	Order Service
}

func LoadConfig() Config {
	viper.SetConfigFile(os.Getenv("CONFIG_PATH"))

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Unable to read config: %v", err)
	}

	var order Service
	if err := viper.UnmarshalKey("Order", &order); err != nil {
		log.Fatalf("Unable to unmarshal Authentication config: %v", err)
	}

	log.Println("Config loaded successfully")

	return Config{
		Order: order,
	}
}
