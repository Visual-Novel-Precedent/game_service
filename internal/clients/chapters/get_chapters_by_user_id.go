package chapters

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"visual_novel/internal/models"
	clients_ "visual_novel/pkg/clients"
)

type GetChaptersRequest struct {
	Id int64 `json:"id"`
}

func GetChaptersByUserId(id int64, log *zerolog.Logger) (*[]models.Chapter, error) {
	client := clients_.NewUniversalHTTPClient()

	reqBody := GetChaptersRequest{
		Id: id,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		log.Err(fmt.Errorf("ошибка при формировании JSON: %w", err))
		return nil, err
	}

	reader := strings.NewReader(string(jsonData))
	req, err := http.NewRequest(http.MethodPost, os.Getenv("GET_CHAPTER_BY_ID"), reader)
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

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Err(fmt.Errorf("ошибка при чтении тела ответа: %w", err))
		return nil, err
	}

	var chapters []models.Chapter
	err = json.Unmarshal(body, &chapters)
	if err != nil {
		log.Err(fmt.Errorf("ошибка при разборе JSON: %w", err))
		return nil, err
	}

	return &chapters, nil
}
