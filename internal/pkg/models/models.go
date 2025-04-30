package models

import "fmt"

type AntiSpoofingResponse struct {
	MelodyMachine float64 `json:"MelodyMachine"`
	MoTheCreator  float64 `json:"mo-thecreator"`
}

func (p *AntiSpoofingResponse) GetWeightedResult() float64 {
	return (p.MelodyMachine*1.0 + p.MoTheCreator*0.6) / (1.0 + 0.6)
}

func (p *AntiSpoofingResponse) GetAnswer() string {
	threshold := 0.7
	if p.GetWeightedResult() > threshold {
		return "Аудио подделано"
	} else {
		return "Аудио настоящее"
	}
}

func (p *AntiSpoofingResponse) ToString() string {
	return fmt.Sprintf(
		"Вероятность спуфинга:\n\tмодель MelodyMachine: %f\n\tмодель mo-thecreator: %f\nСреднее значение: %f\nВзвешенное значение: %f\n\nЗаключение: %v",
		p.MelodyMachine, p.MoTheCreator, (p.MelodyMachine+p.MoTheCreator)/2, p.GetWeightedResult(), p.GetAnswer())
}
