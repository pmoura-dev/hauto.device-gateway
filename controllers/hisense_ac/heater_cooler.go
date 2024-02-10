package hisense_ac

import (
	"errors"
	"fmt"

	"github.com/pmoura-dev/hauto.device-gateway/types"
)

func (c HisenseACController) SetHeatingThresholdTemperature(value int) (types.MQTTData, error) {
	topic, exists := c.actions[types.SetHeatingThresholdTemperature]
	if !exists {
		return types.MQTTData{}, errors.New("action does not exist for this device")
	}

	payload := fmt.Sprintf("%d", value)

	return types.MQTTData{
		Topic:   topic,
		Payload: payload,
	}, nil
}

func (c HisenseACController) SetCoolingThresholdTemperature(value int) (types.MQTTData, error) {
	topic, exists := c.actions[types.SetCoolingThresholdTemperature]
	if !exists {
		return types.MQTTData{}, errors.New("action does not exist for this device")
	}

	payload := fmt.Sprintf("%d", value)

	return types.MQTTData{
		Topic:   topic,
		Payload: payload,
	}, nil
}

func (c HisenseACController) SetHeaterCoolerMode(mode string) (types.MQTTData, error) {
	topic, exists := c.actions[types.SetHeaterCoolerMode]
	if !exists {
		return types.MQTTData{}, errors.New("action does not exist for this device")
	}

	var payload string

	switch mode {
	case types.HeaterCoolerModeHeating:
		payload = "heat"
	case types.HeaterCoolerModeCooling:
		payload = "cool"
	case types.HeaterCoolerModeAutomatic:
	default:
		payload = "auto"
	}

	return types.MQTTData{
		Topic:   topic,
		Payload: payload,
	}, nil
}
