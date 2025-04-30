package models

import "fmt"

type AntiSpoofingResponse struct {
	MelodyMachine float64 `json:"MelodyMachine"`
	DavidCombei   float64 `json:"DavidCombei"`
}

func (p *AntiSpoofingResponse) GetAnswer() bool {
	return ((p.MelodyMachine + p.DavidCombei) / 2) > 0.7
}

func (p *AntiSpoofingResponse) ToString() string {
	return fmt.Sprintf(
		"Вероятность спуфинга:\nмодель MelodyMachine: %s\nмодель DavidCombei: %s\nЗаключение: %v",
		p.MelodyMachine, p.MelodyMachine, p.DavidCombei)
}
