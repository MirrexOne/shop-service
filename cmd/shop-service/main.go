package main

import (
	"github.com/joho/godotenv"
	"shop-service/internal/app"
)

var configPath = "configs/config.yml"

//
//func init() {
//	flag.StringVar(&configPath, "config-path", "configs/config.yml", "path to config file")
//}

func main() {
	//flag.Parse()

	godotenv.Load()
	app.Run(configPath)
}
