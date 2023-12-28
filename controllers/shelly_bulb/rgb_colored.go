package shelly_bulb

import (
	"fmt"
)

func (c ShellyBulbController) SetRGBColor(red, green, blue int) error {
	topic := fmt.Sprintf("shellies/%s/color/0/set", c.NaturalID)
	payload := fmt.Sprintf(`{"red": %d, "green": %d, "blue": %d}`, red, green, blue)

	fmt.Printf("publishing on %s with payload %s\n", topic, payload)
	return nil
}
