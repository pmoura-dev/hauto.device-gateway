package controllers

import (
	"encoding/json"

	"github.com/pmoura-dev/gobroker"
	"github.com/pmoura-dev/hauto.device-gateway/configuration"
	"github.com/pmoura-dev/hauto.device-gateway/controllers/hisense_ac"
	"github.com/pmoura-dev/hauto.device-gateway/controllers/shelly_bulb"
	"github.com/pmoura-dev/hauto.device-gateway/types"
)

const ControllerKey = "controller"

type Switchable interface {
	TurnOn() (types.MQTTData, error)
	TurnOff() (types.MQTTData, error)
}

type RGBColored interface {
	SetRGBColor(red, green, blue int) (types.MQTTData, error)
}

type HeaterCooler interface {
	Heater
	Cooler
	SetHeaterCoolerMode(mode string) (types.MQTTData, error)
}

type Heater interface {
	SetHeatingThresholdTemperature(value int) (types.MQTTData, error)
}

type Cooler interface {
	SetCoolingThresholdTemperature(value int) (types.MQTTData, error)
}

// listeners

type StateListener interface {
	HandleState(rawState []byte) ([]byte, error)
}

type PowerListener interface{}

type EnergyListener interface{}

func GetController(next gobroker.ConsumerHandlerFunc) gobroker.ConsumerHandlerFunc {
	return func(ctx gobroker.ConsumerContext, message gobroker.Message) error {
		var deviceID int

		if ctx.Params["consumer_type"] == "action" {
			var action types.BaseActionMessage

			body := message.GetBody()
			err := json.Unmarshal(body, &action)
			if err != nil {
				return err
			}

			deviceID = action.DeviceID
		} else {
			deviceID = ctx.Params["device_id"].(int)
		}

		mqttConfigurations, _ := configuration.GetDeviceMQTTConfigurations()
		controller := mqttConfigurations[deviceID].Controller
		actions := mqttConfigurations[deviceID].Actions
		listeners := mqttConfigurations[deviceID].Listeners

		switch controller {
		case "shelly_bulb":
			ctx.Params[ControllerKey] = shelly_bulb.NewShellyBulbController(actions, listeners)
		case "hisense_ac":
			ctx.Params[ControllerKey] = hisense_ac.NewHisenseACController(actions, listeners)
		default:
			panic("not implemented")
		}

		return next(ctx, message)
	}
}
