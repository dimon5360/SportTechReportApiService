package main

import (
	"app/main/router"
	"app/main/server"
	"app/main/storage"
	"app/main/utils"
	"fmt"
)

const (
	configPath = "/home/dmitry/Projects/SportTechService/SportTechDockerConfig/"
	serviceEnv = "../config/service.env"
	mongoEnv   = configPath + "mongo.env"
)

func main() {

	utils.Env().Load(serviceEnv, mongoEnv)

	fmt.Println("SportTech report API service v." + utils.Env().Value("SERVICE_VERSION"))

	storage.InitMongo()

	server.InitServer(router.InitRouter(utils.Env().Value("SERVICE_HOST"))).Run()
}
