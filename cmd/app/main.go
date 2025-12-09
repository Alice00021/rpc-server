package main

import (
	"github.com/Alice00021/rpc-server/config"
	"github.com/Alice00021/rpc-server/internal/app"

	"log"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
