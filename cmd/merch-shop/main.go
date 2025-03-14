package main

import (
	"MerchShop/internal/app"
	"flag"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/config.yml", "path to config file")
}

func main() {
	flag.Parse()

	app.Run(configPath)
}
