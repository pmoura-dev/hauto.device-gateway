package shelly_bulb

import (
	"errors"
	"fmt"

	"github.com/pmoura-dev/hauto.device-gateway/types"
)

func (c ShellyBulbController) SetRGBColor(red, green, blue int) (types.MQTTData, error) {
	topic, exists := c.actions[types.SetRGBColor]
	if !exists {
		return types.MQTTData{}, errors.New("action does not exist for this device")
	}
	payload := fmt.Sprintf(`{"red": %d, "green": %d, "blue": %d}`, red, green, blue)

	return types.MQTTData{
		Topic:   topic,
		Payload: payload,
	}, nil
}
