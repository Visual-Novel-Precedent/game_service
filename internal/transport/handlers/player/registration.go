package player_

import (
	"encoding/json"
	"github.com/rs/zerolog"
	"io/ioutil"
	"net/http"
	"visual_novel/internal/services/player"
)

type PlayerRegistrationRequest struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func PlayerRegistrationHandler(log *zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Проверяем, что это POST-запрос
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST requests allowed", http.StatusMethodNotAllowed)
			return
		}

		// Читаем тело запроса
		var req PlayerRegistrationRequest
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

		id, err := player.Registration(req.Email, req.Name, req.Password, log)

		if err != nil {
			http.Error(w, "fail to register player", http.StatusInternalServerError)
		}

		// Формируем ответ
		response := map[string]interface{}{
			"id": *id,
		}

		// Отправляем ответ клиенту
		json.NewEncoder(w).Encode(response)
	}
}
