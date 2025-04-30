package anti_spoof_client

import (
	"encoding/json"
	"fmt"
	"github.com/cockroachdb/errors"
	"io"
	"net/http"
)

type Predict struct {
	//модель 1 например
	res1 string
	//модель 2 например
	res2 string
	//Итоговый ответ, считаться может по формуле
	ans bool
}

func (p *Predict) ToString() string {
	return fmt.Sprintf("model1: %s\nmodel2: %s\nresult:%v", p.res1, p.res2, p.ans)
}

type Service struct {
	client http.Client
	url    string
}

// Принимает на вход путь до wav файла и отправляет grpc запрос к external-api api
func (s *Service) GetPredict(path string) (*Predict, error) {
	//s.client
	// Создаём новый http-клиент и запрос
	req, err := http.NewRequest("GET", s.url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "http.NewRequest")
	}

	// Добавляем строковую переменную в заголовке (например, X-My-Variable)
	req.Header.Set("X-Audio-Path", path)

	// Выполняем запрос
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "client.Do")
	}
	defer resp.Body.Close()

	// Читаем тело ответа
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "io.ReadAll")
	}

	// Выводим полученный ответ как текст
	fmt.Println("Ответ сервера (raw):")
	fmt.Println(string(body))

	// Если сервер возвращает JSON, парсим его
	var prediction Predict
	err = json.Unmarshal(body, &prediction)
	if err != nil {
		return nil, errors.Wrap(err, "json.Unmarshal")
	}

	return &prediction, nil
}
