package shelly_bulb

type ShellyBulbController struct {
	actions   map[string]string
	listeners map[string]string
}

func NewShellyBulbController(actions, listeners map[string]string) ShellyBulbController {
	return ShellyBulbController{
		actions:   actions,
		listeners: listeners,
	}
}
