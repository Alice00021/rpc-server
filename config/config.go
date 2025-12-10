package config

import (
	"fmt"
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"log"
)

type (
	// Config -.
	Config struct {
		App  App
		HTTP HTTP
		Log  Log
	}

	// App -.
	App struct {
		Name    string `env:"APP_NAME,required"`
		Version string `env:"APP_VERSION,required"`
	}

	// HTTP -.
	HTTP struct {
		Port string `env:"HTTP_PORT,required"`
	}

	// Log -.
	Log struct {
		Level    string `env:"LOG_LEVEL,required"`
		FileName string `env:"LOG_FILE_NAME,required"`
	}
)

func NewConfig() (*Config, error) {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("Could not loading .env file")
	}

	// Parse environment variables into structs
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	return cfg, nil
}
