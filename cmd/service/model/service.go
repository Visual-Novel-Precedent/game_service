package model

import (
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
	"visual_novel/pkg/config"
	"visual_novel/pkg/db"
	"visual_novel/pkg/log"
	"visual_novel/pkg/router"
)

type Service struct {
	Log    *zerolog.Logger
	Router *mux.Router
	DB     *gorm.DB
	Config *config.Config
}

func NewService() *Service {
	logger := log.NewLogger()

	router := router.NewRouter()

	db, err := db.InitDB()

	if err != nil {
		logger.Error().Msg("error to get db")
	}

	config := config.NewConfig()

	return &Service{
		Log:    logger,
		Router: router,
		DB:     db,
		Config: config,
	}
}
