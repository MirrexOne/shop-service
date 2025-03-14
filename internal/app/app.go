package app

import (
	"MerchShop/internal/config"
	"MerchShop/internal/database"
	"log"
)

func Run(configPath string) {
	// Configuration
	cfg, err := config.NewConfig(configPath)
	if err != nil {
		log.Fatal("Error occurred while creating new config")
	}

	log.Print(cfg)

	// Initialize database
	db, err := database.New(cfg.DB.URL)
	if err != nil {
		log.Fatalf("Error occurred while initializing database: %v", db)
	}

	log.Print("Database successfully initialized")

	defer db.Close()
}
