package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	IsDebug  bool   `json:"is_debug"`
	ApiToken string `json:"api_token"`
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Panic("Error loading .env file")
	}

	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config
	cfg.ApiToken = os.Getenv("API_TOKEN")
	err = cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("an error occurred while reading config: %s", configPath)
	}

	return &cfg
}
