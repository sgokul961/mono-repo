package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/sgokul961/GO-PROJECT/pkg/config"
	"github.com/sgokul961/GO-PROJECT/pkg/di"
)

func main() {
	//	gin.SetMode(gin.ReleaseMode)

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config, configErr := config.LoadConfig()

	if configErr != nil {

		log.Fatal("cannot load config:", configErr)

	}

	server, diErr := di.InitializeApi(config)

	if diErr != nil {
		log.Fatal("cannot start server:", diErr)
	} else {
		server.Start()
	}

}
