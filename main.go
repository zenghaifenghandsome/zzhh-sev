package main

import (
	"z-web-sev/config"
	"z-web-sev/model"
	"z-web-sev/router"
)

func main() {
	config.InitConfig()
	model.InitDb()
	//model.InitMongoDb()
	router.InitRouter()
}
