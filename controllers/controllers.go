package controllers

import (
	"encoding/json"

	"github.com/pmoura-dev/gobroker"
	"github.com/pmoura-dev/hauto.device-gateway/controllers/shelly_bulb"
	"github.com/pmoura-dev/hauto.device-gateway/types"
)

const ControllerKey = "controller"

type Switchable interface {
	TurnOn() types.MQTTData
	TurnOff() types.MQTTData
}

type RGBColored interface {
	SetRGBColor(red, green, blue int) types.MQTTData
}

func GetController(next gobroker.ConsumerHandlerFunc) gobroker.ConsumerHandlerFunc {
	return func(ctx gobroker.ConsumerContext, message gobroker.Message) error {
		var action types.BaseActionMessage

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
