package clients

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const (
	DefaultTimeout    = 30 * time.Second
	DefaultTimeoutStr = "30"
)

type UniversalHTTPClient struct {
	client *http.Client
}

func NewUniversalHTTPClient() *UniversalHTTPClient {
	timeoutStr := os.Getenv("HTTP_CLIENT_TIMEOUT")

	// Если переменная окружения не установлена, используем значение по умолчанию
	if timeoutStr == "" {
		timeoutStr = DefaultTimeoutStr
	}

	// Преобразуем строку в время.Duration
	timeout, err := time.ParseDuration(timeoutStr + "s")

	if err != nil {
		fmt.Printf("Ошибка при парсинге времени: %v\n", err)
		// Используем значение по умолчанию
		timeout = DefaultTimeout
	}

	return &UniversalHTTPClient{
		client: &http.Client{
			Timeout: timeout,
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return nil // не переопределяем проверку перенаправлений
			},
		},
	}
}

func (c *UniversalHTTPClient) Do(req *http.Request) (*http.Response, error) {
	return c.client.Do(req)
}
