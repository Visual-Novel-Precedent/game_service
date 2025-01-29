package player

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog"
	"net/http"
	"os"
	"strings"
	clients_ "visual_novel/pkg/clients"
)

type PlayerChapterProgressRequest struct {
	Id        int64 `json:"id"`
	ChapterId int64 `json:"chapter_id"`
	NodeId    int64 `json:"node_id"`
	EndFlag   bool  `json:"end_flag"`
}

func ChangePlayerProgress(id int64, chapterId int64, nodeId int64, endFlag bool, log *zerolog.Logger) error {
	client := clients_.NewUniversalHTTPClient()

	reqBody := PlayerChapterProgressRequest{
		Id:        id,
		ChapterId: chapterId,
		NodeId:    nodeId,
		EndFlag:   endFlag,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		log.Err(fmt.Errorf("ошибка при формировании JSON: %w", err))
		return err
	}

	reader := strings.NewReader(string(jsonData))
	req, err := http.NewRequest(http.MethodPost, os.Getenv("UPDATE_CHAPTER_PROGRESS"), reader)
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
