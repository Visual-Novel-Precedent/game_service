package player

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog"
	"net/http"
	"strings"
	clients_ "visual_novel/pkg/clients "
)

type ChangePlayerRequest struct {
	Id            int64  `json:"id"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	Phone         string `json:"phone"`
	Password      string `json:"password"`
	SoundSettings int    `json:"sound_settings"`
}

func ChangePlayer(email string, id int64, name string, phone string, password string, soundSettings int, log *zerolog.Logger) error {
	client := clients_.NewUniversalHTTPClient()

	reqBody := ChangePlayerRequest{
		Email:         email,
		Name:          name,
		Id:            id,
		Phone:         phone,
		Password:      password,
		SoundSettings: soundSettings,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		log.Err(fmt.Errorf("ошибка при формировании JSON: %w", err))
		return err
	}

	reader := strings.NewReader(string(jsonData))
	req, err := http.NewRequest(http.MethodPost, "http://your-api-url/admin-authorization", reader)
	if err != nil {
		log.Err(fmt.Errorf("Ошибка при создании запроса: %w", err))
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Err(fmt.Errorf("Ошибка при отправке запроса: %w", err))
		return err
	}
	defer resp.Body.Close()

	return nil
}
