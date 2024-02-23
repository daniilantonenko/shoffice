package main

import (
	"log"
	"shoffice/internal/app"
)

func main() {
	cfg, err := app.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	app.Run(cfg)
}
