package player

import (
	"github.com/rs/zerolog"
	"visual_novel/internal/clients/player"
)

func ChangePlayer(
	id int64,
	name string,
	email string,
	phone string,
	password string,
	soundSettings int,
	log *zerolog.Logger,
) error {
	return player.ChangePlayer(email, id, name, phone, password, soundSettings, log)
}
