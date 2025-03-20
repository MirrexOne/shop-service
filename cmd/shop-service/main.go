package main

import (
	"shop-service/internal/app"
)

var configPath = "configs/config.yml"

//
//func init() {
//	flag.StringVar(&configPath, "config-path", "configs/config.yml", "path to config file")
//}

func main() {
	//flag.Parse()

	app.Run(configPath)
}
