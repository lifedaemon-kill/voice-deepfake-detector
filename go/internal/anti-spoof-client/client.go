package anti_spoof_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/cockroachdb/errors"
	"main/internal/pkg/config"
	"main/internal/pkg/models"
	"net/http"
)

type Client struct {
	Host              string
	AntiSpoofEndpoint string
}

func NewClient(config config.AppConfig) *Client {
	return &Client{
		Host:              config.ASHost,
		AntiSpoofEndpoint: config.AASEndpoint,
	}
}
func (c *Client) SendRequest(filePath string) (*models.AntiSpoofingResponse, error) {
	requestBody, err := json.Marshal(map[string]string{"file_path": filePath})
	if err != nil {
		return nil, fmt.Errorf("ошибка при маршализации JSON: %v", err)
	}

	resp, err := http.Post(c.Host+c.AntiSpoofEndpoint, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, errors.Wrap(err, "ошибка при отправке запроса")
	}
	defer resp.Body.Close()

	// Проверяем статус ответа
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("неожиданный статус ответа: " + resp.Status + "; body: " + string(requestBody) + "; endpoint: " + c.Host + c.AntiSpoofEndpoint + ";")
	}

	// Декодируем JSON-ответ
	var response models.AntiSpoofingResponse
	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("ошибка при декодировании JSON: %v", err)
	}

	return &response, nil
}
