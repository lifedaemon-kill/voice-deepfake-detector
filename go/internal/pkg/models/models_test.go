package models

import (
	"fmt"
	"testing"
)

func TestToString(t *testing.T) {
	resp := AntiSpoofingResponse{
		MelodyMachine: 0.999123,
		MoTheCreator:  0.9123,
	}

	fmt.Println(resp.ToString())
}
