package shelly_bulb

import (
	"fmt"

	"github.com/pmoura-dev/hauto.device-gateway/types"
)

func (c ShellyBulbController) TurnOn() types.MQTTData {
	topic := fmt.Sprintf("shellies.%s.color.0.command", c.NaturalID)
	payload := "on"

	return types.MQTTData{
		Topic:   topic,
		Payload: payload,
	}
}

func (c ShellyBulbController) TurnOff() types.MQTTData {
	topic := fmt.Sprintf("shellies.%s.color.0.command", c.NaturalID)
	payload := "off"

	return types.MQTTData{
		Topic:   topic,
		Payload: payload,
	}
}
