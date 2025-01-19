package main

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"visual_novel/cmd/service/model"
)

// init is invoked before main()
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

const (
	PORT = "APP_PORT"
)

func main() {
	service := run()

	port := os.Getenv(PORT)
	if port == "" {
		port = "8080" // Default port if not set
	}

	// Запуск сервера
	http.ListenAndServe(port, service.Router)
}

func run() *model.Service {
	service := model.NewService()

	service.Log.Info().Msg("service is created ")

	return service
}
