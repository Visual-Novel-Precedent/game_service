package player

import (
	"github.com/rs/zerolog"
	"visual_novel/internal/clients/player"
)

func UpdateChapterProgress(id int64, chapterId int64, nodeId int64, endFlag bool, log *zerolog.Logger) error {

	return player.ChangePlayerProgress(id, chapterId, nodeId, endFlag, log)
}
