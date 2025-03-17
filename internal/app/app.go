package app

import (
	log "github.com/sirupsen/logrus"
	"shop-service/internal/config"
	v1 "shop-service/internal/controller/http/v1"
	"shop-service/internal/database"
	"shop-service/internal/repo"
	"shop-service/internal/server"
	"shop-service/internal/service"
)

func Run(configPath string) {
	notify := make(chan int)

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
	defer db.Close()

	log.Info("Database successfully started")

	// Initialize repositories
	repositories := repo.NewRepositories(db)

	// Services dependencies
	deps := service.ServicesDependencies{
		Repos: repositories,
	}
	services := service.NewServices(deps)

	log.Info("Initialize handlers and routes")

	r := v1.New()

	v1.NewRouter(r, services)

	log.Info("Starting http server...")
	log.Debugf("Server port: %s", cfg.Server.Port)
	server.NewServer(r.Mux, server.Port(cfg.Server.Port))
	<-notify
}
