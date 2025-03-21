package app

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"shop-service/internal/config"
	v1 "shop-service/internal/controller/http/v1"
	"shop-service/internal/database"
	"shop-service/internal/repo"
	"shop-service/internal/server"
	"shop-service/internal/service"
	"shop-service/pkg/hasher"
	"syscall"
)

func Run(configPath string) {

	// Configuration
	cfg, err := config.NewConfig(configPath)
	if err != nil {
		log.Fatalf("Error occurred while creating new config: %s", err)
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
		Repos:  repositories,
		Hasher: hasher.NewSHA1Hasher(cfg.Hasher.Salt),
	}
	services := service.NewServices(deps)

	log.Info("Initialize handlers and routes")

	// Initialize routes
	r := v1.New()
	v1.NewRouter(r, services)

	log.Info("Starting http server...")
	log.Debugf("Server port: %s", cfg.Server.Port)
	srv := server.NewServer(r.Mux, server.Port(cfg.Server.Port))

	// Waiting signal
	log.Info("Configuring graceful shutdown...")
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Info("app signal: " + s.String())
	case err = <-srv.Notify():
		log.Error(fmt.Errorf("server notify: %w", err))
	}

	// Graceful shutdown
	log.Info("Shutting down...")
	err = srv.Shutdown()
	if err != nil {
		log.Error(fmt.Errorf("srv.Shutdown: %w", err))
	}
}
