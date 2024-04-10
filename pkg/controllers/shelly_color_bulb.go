package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/crazy3lf/colorconv"
	"github.com/pmoura-dev/hauto.device-gateway/pkg/states"
	"github.com/pmoura-dev/hauto.device-gateway/pkg/types"
)

const (
	ShellyColorBulbControllerName = "shelly_color_bulb"
)

type ShellyColorBulbController struct {
	BaseController
}

func NewShellyColorBulbController(naturalID string) ShellyColorBulbController {
	return ShellyColorBulbController{
		BaseController{naturalID: naturalID},
	}
}

func (c ShellyColorBulbController) GetStateTopic() string {
	const topicFormat = "shellies.%s.color.0.status"

	return fmt.Sprintf(topicFormat, c.naturalID)
}

func (c ShellyColorBulbController) HandleState(rawState []byte, _ string) (states.State, error) {
	var parsedState shellyColorBulbRawState

	if err := json.Unmarshal(rawState, &parsedState); err != nil {
		return nil, err
	}

	status := types.OfflineStatus
	if parsedState.IsOn {
		status = types.OnlineStatus
	}

	mode := types.ColorLightColorMode
	if parsedState.Mode == "white" {
		mode = types.ColorLightWhiteMode
	}

	brightness := parsedState.Brightness
	temperature := parsedState.Temp

	h, s, l := colorconv.RGBToHSL(parsedState.Red, parsedState.Green, parsedState.Blue)
	return &states.ColorLightState{
		Status: status,
		Mode:   mode,
		Color: types.HSLAColor{
			Hue:        int(h),
			Saturation: int(s * 100),
			Lightness:  int(l * 100),
		},
		Brightness:  brightness,
		Temperature: temperature,
	}, nil
}

type shellyColorBulbRawState struct {
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

func (c ShellyColorBulbController) TurnOn() MQTTData {
	const topicFormat = "shellies.%s.color.0.command"

	return MQTTData{
		Topic:   fmt.Sprintf(topicFormat, c.naturalID),
		Payload: "on",
	}
}

func (c ShellyColorBulbController) TurnOff() MQTTData {
	const topicFormat = "shellies.%s.color.0.command"

	return MQTTData{
		Topic:   fmt.Sprintf(topicFormat, c.naturalID),
		Payload: "off",
	}
}
