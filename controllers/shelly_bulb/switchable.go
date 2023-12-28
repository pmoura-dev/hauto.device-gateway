package shelly_bulb

import (
	"fmt"
)

func (c ShellyBulbController) TurnOn() error {
	topic := fmt.Sprintf("shellies/%s/color/0/command", c.NaturalID)
	payload := "on"

	fmt.Printf("publishing on %s with payload %s\n", topic, payload)

	// TODO: publish mqtt message
	return nil
}

func (c ShellyBulbController) TurnOff() error {
	topic := fmt.Sprintf("shellies/%s/color/0/command", c.NaturalID)
	payload := "off"

	fmt.Printf("publishing on %s with payload %s\n", topic, payload)
	return nil
}
