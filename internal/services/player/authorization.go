package player

import (
	"github.com/rs/zerolog"
	chapters2 "visual_novel/internal/clients/chapters"
	characters2 "visual_novel/internal/clients/characters "
	"visual_novel/internal/clients/player"
	"visual_novel/internal/models"
)

func AuthorizationPlayer(email string, password string, log *zerolog.Logger) (*models.Player, *[]models.Chapter, *[]models.Character, error) {
	user, err := player.AuthorizationPlayer(email, password, log)

	if err != nil {
		return nil, nil, nil, err
	}

	chapters, err := chapters2.GetChaptersByUserId(user.Id, log)

	if err != nil {
		return nil, nil, nil, err
	}

	characters, err := characters2.GetCharacters(log)

	if err != nil {
		return nil, nil, nil, err
	}

	return user, chapters, characters, nil
}
