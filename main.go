package main

import (
	"z-web-sev/model"
	"z-web-sev/router"
)

func main() {
	model.InitDb()
	//model.InitMongoDb()
	router.InitRouter()
}
