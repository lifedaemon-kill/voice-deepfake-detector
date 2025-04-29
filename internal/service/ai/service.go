package ai

import "fmt"

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
}

// Принимает на вход путь до wav файла и отправляет grpc запрос к python api
func (s *Service) GetPredict(path string) *Predict {
	return nil
}
