package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
)

type Config struct {
	Env             string `yaml:".env" .env-default:"local"`
	StoragePath     string `yaml:"storage_path" .env-default:"./storage/storage.db"`
	StorageAddress  string `yaml:"storage_address" .env-default:"./storage/storage.db"`
	StorageUser     string `yaml:"storage_user" .env-default:"user"`
	StoragePassword string `yaml:"storage_password" .env-default:"password"`
	HTTPServer      `yaml:"http_server"`
}

type HTTPServer struct {
	Address     string `yaml:"address" .env-default:"localhost:8080"`
	Timeout     string `yaml:"timeout" .env-default:"4s"`
	IdleTimeout string `yaml:"idle_timeout" .env-default:"60s"`
}

func MustLoad() *Config {
	//configPath := os.Getenv("CONFIG_PATH")
	configPath := "./config/local.yaml"
	if configPath == "" {
		log.Fatal("CONFIG_PATH environment variable not set")
	}

	// check if file not exist
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Config file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("Config file does not exist: %s", configPath)
	}

	return &cfg
}
