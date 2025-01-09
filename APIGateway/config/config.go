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
	Gateway        Service
	Authentication Service
	Order          Service
	Status         Service
	Payment        Service
	History        Service
}

func LoadConfig() Config {
	viper.SetConfigFile(os.Getenv("CONFIG_PATH"))

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Unable to read config: %v", err)
	}

	var gateway Service
	if err := viper.UnmarshalKey("Gateway", &gateway); err != nil {
		log.Fatalf("Unable to unmarshal Gateway config: %v", err)
	}

	var auth Service
	if err := viper.UnmarshalKey("Authentication", &auth); err != nil {
		log.Fatalf("Unable to unmarshal Authentication config: %v", err)
	}

	var order Service
	if err := viper.UnmarshalKey("Order", &order); err != nil {
		log.Fatalf("Unable to unmarshal Order config: %v", err)
	}

	var status Service
	if err := viper.UnmarshalKey("Status", &status); err != nil {
		log.Fatalf("Unable to unmarshal Status config: %v", err)
	}

	var payment Service
	if err := viper.UnmarshalKey("Payment", &payment); err != nil {
		log.Fatalf("Unable to unmarshal Payment config: %v", err)
	}

	var history Service
	if err := viper.UnmarshalKey("History", &history); err != nil {
		log.Fatalf("Unable to unmarshal History config: %v", err)
	}

	log.Println("Config loaded successfully")

	return Config{
		Gateway:        gateway,
		Authentication: auth,
		Order:          order,
		Status:         status,
		Payment:        payment,
		History:        history,
	}
}
