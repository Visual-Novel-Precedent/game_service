package player

import (
	"github.com/rs/zerolog"
	"visual_novel/internal/clients/player"
)

func Registration(email string, name string, password string, log *zerolog.Logger) (*int64, error) {
	return player.PlayerRegistration(email, password, name, log)
}
