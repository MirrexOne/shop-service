package app

import (
	log "github.com/sirupsen/logrus"
	"shop-service/internal/config"
	"shop-service/internal/database"
)

func Run(configPath string) {
	// Configuration
	cfg, err := config.NewConfig(configPath)
	if err != nil {
		log.Fatal("Error occurred while creating new config")
	}

	log.Info("Config successfully parsed")

	// Logger
	SetLogrus(cfg.Log.Level)

	// Initialize database
	db, err := database.New(cfg.DB.URL)
	if err != nil {
		log.Fatalf("Error occurred while initializing database: %v", db)
	}

	log.Info("Database successfully started")

	defer db.Close()

	log.Info("Starting http server...")
	log.Debugf("Server port: %s", cfg.Server.Port)
}
