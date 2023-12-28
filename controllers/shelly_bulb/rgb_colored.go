package shelly_bulb

import (
	"fmt"

	"github.com/pmoura-dev/hauto.device-gateway/types"
)

func (c ShellyBulbController) SetRGBColor(red, green, blue int) types.MQTTData {
	topic := fmt.Sprintf("shellies.%s.color.0.set", c.NaturalID)
	payload := fmt.Sprintf(`{"red": %d, "green": %d, "blue": %d}`, red, green, blue)

	return types.MQTTData{
		Topic:   topic,
		Payload: payload,
	}
}
