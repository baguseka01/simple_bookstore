package main

import (
	"log"
	"os"

	"github.com/baguseka01/simple_bookstore/config"
	"github.com/baguseka01/simple_bookstore/routes"
	"github.com/joho/godotenv"
)

type AppConfig struct {
	AppName string
	AppEnv  string
	AppPort string
}

func main() {
	var appConfig = AppConfig{}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error on loading .env file")
	}

	// Configuration App
	appConfig.AppName = os.Getenv("APP_NAME")
	appConfig.AppEnv = os.Getenv("APP_ENV")
	appConfig.AppPort = os.Getenv("APP_PORT")

	// Initialize Database
	config.Connect()
	config.Migration()

	// Initialize Router
	router := routes.Router()
	router.Start(":" + appConfig.AppPort)
}
