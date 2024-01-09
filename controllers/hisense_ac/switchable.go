package hisenseac

import (
	"fmt"

	"github.com/pmoura-dev/hauto.device-gateway/types"
)

func (c *HisenseACController) TurnOn() types.MQTTData {
	topic := fmt.Sprintf("hisense_ac.%s.t_power.command", c.NaturalID)
	payload := "on"

	return types.MQTTData{
		Topic:   topic,
		Payload: payload,
	}
}

func (c *HisenseACController) TurnOff() types.MQTTData {
	topic := fmt.Sprintf("hisense_ac.%s.t_power.command", c.NaturalID)
	payload := "off"

	return types.MQTTData{
		Topic:   topic,
		Payload: payload,
	}
}
