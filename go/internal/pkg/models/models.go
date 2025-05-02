package models

import "fmt"

type AntiSpoofingResponse struct {
	MelodyMachine float64 `json:"MelodyMachine"`
	MoTheCreator  float64 `json:"mo-thecreator"`
}

func (p *AntiSpoofingResponse) WeightedAverage() float64 {
	//Взвешенное среднее нужно для того, чтобы менее точная модель имела меньшее влияние на результат
	return (p.MelodyMachine*1.0 + p.MoTheCreator*0.6) / (1.0 + 0.6)
}

func (p *AntiSpoofingResponse) Predict() string {
	threshold := 0.7
	if p.WeightedAverage() > threshold {
		return "Аудио подделано или модифицировано"
	} else {
		return "Аудио настоящее"
	}
}

func (p *AntiSpoofingResponse) ToString(filename string) string {
	return fmt.Sprintf(
		filename+"\nВероятность спуфинга:\n\tm1: %f\n\tm2: %f\n Avg: %f\nwAvg: %f\n\nЗаключение: %v",
		p.MelodyMachine, p.MoTheCreator, (p.MelodyMachine+p.MoTheCreator)/2, p.WeightedAverage(), p.Predict())
}
