package player_

import (
	"encoding/json"
	"github.com/rs/zerolog"
	"io/ioutil"
	"net/http"
	player "visual_novel/internal/services/player"
)

type PlayerAuthorisationRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func PlayerAuthorisationHandler(log *zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Проверяем, что это POST-запрос
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST requests allowed", http.StatusMethodNotAllowed)
			return
		}

		// Читаем тело запроса
		var req PlayerAuthorisationRequest
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

		// Здесь должна быть логика получения данных пользователя
		// Например, из базы данных:
		player, chapters, characters, err := player.AuthorizationPlayer(req.Email, req.Password, log)

		// Формируем ответ
		response := map[string]interface{}{
			"player":     *player,
			"chapters":   *chapters,
			"characters": *characters,
		}

		// Отправляем ответ клиенту
		json.NewEncoder(w).Encode(response)
	}
}
