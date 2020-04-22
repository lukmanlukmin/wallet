package main

import (
	"config"
	"controller"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	if godotenv.Load() != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	startApp()
}

func startApp() {
	router := config.SetupRouter()
	controller.LoadRouter(router)
	serverHost := os.Getenv("SERVER_ADDRESS")
	serverPort := os.Getenv("SERVER_PORT")
	serverString := fmt.Sprintf("%s:%s", serverHost, serverPort)
	fmt.Println(serverString)
	router.Run(serverString)
}
