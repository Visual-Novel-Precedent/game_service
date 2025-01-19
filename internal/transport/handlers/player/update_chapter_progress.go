package player_

import (
	"encoding/json"
	"github.com/rs/zerolog"
	"io/ioutil"
	"net/http"
	"visual_novel/internal/services/player"
)

type PlayerChapterProgressRequest struct {
	Id        int64 `json:"id"`
	ChapterId int64 `json:"chapter_id"`
	NodeId    int64 `json:"node_id"`
	EndFlag   bool  `json:"end_flag"`
}

func PlayerChapterProgressHandler(log *zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Проверяем, что это POST-запрос
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST requests allowed", http.StatusMethodNotAllowed)
			return
		}

		// Читаем тело запроса
		var req PlayerChapterProgressRequest
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusInternalServerError)
			return
		}

		// Разбираем JSON
		err = json.Unmarshal(body, &req)
		if err != nil {
			http.Error(w, "Invalid JSON format", http.StatusBadRequest)
			return
		}

		err = player.UpdateChapterProgress(req.Id, req.ChapterId, req.NodeId, req.EndFlag, log)

		if err != nil {
			http.Error(w, "Fail to update chapter progress", http.StatusInternalServerError)
			return
		}
	}
}
