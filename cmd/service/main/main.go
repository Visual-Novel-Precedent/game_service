package main

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"visual_novel/cmd/service/model"
	"visual_novel/internal/transport/handlers/node"
	player_ "visual_novel/internal/transport/handlers/player"
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

	service.Router.HandleFunc("/get_nodes", func(w http.ResponseWriter, r *http.Request) {
		handler := node.GetNodeByChapterIdHandler(service.Log)
		handler.ServeHTTP(w, r)
	})

	service.Router.HandleFunc("/player_registration", func(w http.ResponseWriter, r *http.Request) {
		handler := player_.PlayerRegistrationHandler(service.Log)
		handler.ServeHTTP(w, r)
	})
	service.Router.HandleFunc("/player_chapter_progress_change", func(w http.ResponseWriter, r *http.Request) {
		handler := player_.PlayerChapterProgressHandler(service.Log)
		handler.ServeHTTP(w, r)
	})
	service.Router.HandleFunc("/player_authorization", func(w http.ResponseWriter, r *http.Request) {
		handler := player_.PlayerAuthorisationHandler(service.Log)
		handler.ServeHTTP(w, r)
	})
	service.Router.HandleFunc("/change_player", func(w http.ResponseWriter, r *http.Request) {
		handler := player_.ChangePlayerHandler(service.Log)
		handler.ServeHTTP(w, r)
	})

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
