package main

import (
	cfg "MerchShop/internal/config"
	"flag"
	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/config.yml", "path to config file")
}

func main() {
	flag.Parse()

	config, err := cfg.NewConfig(configPath)
	if err != nil {
		log.Fatal("Error occurred while creating new config")
	}

	log.Print(config)
}
