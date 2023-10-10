package main

import (
	"log"
	"github.com/dexxp/L0/internal/app"
	"github.com/dexxp/L0/config"
)

func main() {
	cfg, err := config.NewConfig()

	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	app.Run(cfg)
}