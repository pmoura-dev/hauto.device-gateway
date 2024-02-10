package shelly_bulb

import (
	"encoding/json"

	"github.com/crazy3lf/colorconv"
	"github.com/pmoura-dev/hauto.device-gateway/types"
	"github.com/shopspring/decimal"
)

func (c ShellyBulbController) HandleState(rawState []byte) ([]byte, error) {
	var state shellyBulbState

	err := json.Unmarshal(rawState, &state)
	if err != nil {
		return nil, err
	}

	status := "offline"
	if state.IsOn {
		status = "online"
	}

	hue, saturation, lightness := colorconv.RGBToHSL(state.Red, state.Green, state.Blue)

	convertedState := types.LightBulbState{
		Status:      status,
		Mode:        state.Mode,
		Hue:         decimal.NewFromFloat(hue).Round(1),
		Saturation:  decimal.NewFromFloat(saturation * 100).Round(1),
		Lightness:   decimal.NewFromFloat(lightness * 100).Round(1),
		Brightness:  state.Gain,
		Temperature: state.Temp,
	}

	return json.Marshal(convertedState)
}

type shellyBulbState struct {
	IsOn           bool   `json:"ison"`
	HasTimer       bool   `json:"has_timer"`
	TimerStarted   int    `json:"timer_started"`
	TimerDuration  int    `json:"timer_duration"`
	TimerRemaining int    `json:"timer_remaining"`
	Mode           string `json:"mode"`
	Red            uint8  `json:"red"`
	Green          uint8  `json:"green"`
	Blue           uint8  `json:"blue"`
	White          int    `json:"white"`
	Gain           int    `json:"gain"`
	Temp           int    `json:"temp"`
	Brightness     int    `json:"brightness"`
	Effect         int    `json:"effect"`
}
