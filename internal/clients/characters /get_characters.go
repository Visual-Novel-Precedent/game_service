package characters_

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog"
	"io"
	"net/http"
	"os"
	"visual_novel/internal/models"
	clients_ "visual_novel/pkg/clients"
)

func GetCharacters(log *zerolog.Logger) (*[]models.Character, error) {
	client := clients_.NewUniversalHTTPClient()

	// Create the request URL with the chapter ID as a query parameter
	reqURL := fmt.Sprintf(os.Getenv("GET_CHARACTERS"))

	req, err := http.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		log.Err(fmt.Errorf("Ошибка при создании запроса: %w", err))
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Err(fmt.Errorf("Ошибка при отправке запроса: %w", err))
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Err(fmt.Errorf("ошибка при чтении тела ответа: %w", err))
		return nil, err
	}

	var characters []models.Character
	err = json.Unmarshal(body, &characters)
	if err != nil {
		log.Err(fmt.Errorf("ошибка при разборе JSON: %w", err))
		return nil, err
	}

	return &characters, nil
}
