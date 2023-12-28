package controllers

import (
	"encoding/json"

	"github.com/pmoura-dev/gobroker"
	"github.com/pmoura-dev/hauto.device-gateway/api"
	"github.com/pmoura-dev/hauto.device-gateway/controllers/shelly_bulb"
)

const ControllerKey = "controller"

type Switchable interface {
	TurnOn() error
	TurnOff() error
}

type RGBColored interface {
	SetRGBColor(red, green, blue int) error
}

func GetController(next gobroker.ConsumerHandlerFunc) gobroker.ConsumerHandlerFunc {
	return func(ctx gobroker.ConsumerContext, message gobroker.Message) error {
		var action api.BaseActionMessage

		body := message.GetBody()
		err := json.Unmarshal(body, &action)
		if err != nil {
			return err
		}

		switch action.Controller {
		case "shelly_bulb":
			ctx.Params[ControllerKey] = shelly_bulb.ShellyBulbController{NaturalID: action.NaturalID}
		case "hisense_ac":
			panic("not implemented")
		}

		return next(ctx, message)
	}
}
