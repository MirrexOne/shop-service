package app

import (
	"MerchShop/internal/config"
	"log"
)

func Run(configPath string) {
	// Configuration
	cfg, err := config.NewConfig(configPath)
	if err != nil {
		log.Fatal("Error occurred while creating new config")
	}

	log.Print(cfg)
}
