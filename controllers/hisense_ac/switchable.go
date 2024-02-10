package hisense_ac

import (
	"errors"

	"github.com/pmoura-dev/hauto.device-gateway/types"
)

func (c HisenseACController) TurnOn() (types.MQTTData, error) {
	topic, exists := c.actions[types.TurnOnAction]
	if !exists {
		return types.MQTTData{}, errors.New("action does not exist for this device")
	}

	payload := "on"

	return types.MQTTData{
		Topic:   topic,
		Payload: payload,
	}, nil
}

func (c HisenseACController) TurnOff() (types.MQTTData, error) {
	topic, exists := c.actions[types.TurnOffAction]
	if !exists {
		return types.MQTTData{}, errors.New("action does not exist for this device")
	}

	payload := "off"

	return types.MQTTData{
		Topic:   topic,
		Payload: payload,
	}, nil
}
