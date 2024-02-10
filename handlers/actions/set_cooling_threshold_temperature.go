package actions

import (
	"encoding/json"

	"github.com/pmoura-dev/gobroker"
	"github.com/pmoura-dev/hauto.device-gateway/controllers"
	"github.com/pmoura-dev/hauto.device-gateway/types"
)

func SetCoolingThresholdTemperature(ctx gobroker.ConsumerContext, message gobroker.Message) error {
	var action types.SetTemperatureThresholdMessage
	err := json.Unmarshal(message.GetBody(), &action)
	if err != nil {
		return err
	}

	controller := ctx.Params[controllers.ControllerKey].(controllers.Cooler)

	value := action.Value
	data, err := controller.SetCoolingThresholdTemperature(value)
	if err != nil {
		return err
	}

	return ctx.Publisher.Publish([]byte(data.Payload), data.Topic, map[string]any{
		"exchange": "amq.topic",
	})
}
