package player

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog"
	"io/ioutil"
	"net/http"
	"strings"
	clients_ "visual_novel/pkg/clients "
)

type PlayerRegistrationRequest struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func PlayerRegistration(email string, password string, name string, log *zerolog.Logger) (*int64, error) {
	client := clients_.NewUniversalHTTPClient()

	reqBody := PlayerRegistrationRequest{
		Email:    email,
		Name:     name,
		Password: password,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		log.Err(fmt.Errorf("ошибка при формировании JSON: %w", err))
		return nil, err
	}

	reader := strings.NewReader(string(jsonData))
	req, err := http.NewRequest(http.MethodPost, "http://your-api-url/admin-authorization", reader)
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

	if resp.StatusCode == http.StatusForbidden {
		return nil, fmt.Errorf("invalid password")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Err(fmt.Errorf("ошибка при чтении тела ответа: %w", err))
		return nil, err
	}

	var id int64
	err = json.Unmarshal(body, &id)
	if err != nil {
		log.Err(fmt.Errorf("ошибка при разборе JSON: %w", err))
		return nil, err
	}

	return &id, nil
}
