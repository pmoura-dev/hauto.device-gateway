package actions

import (
	"encoding/json"

	"github.com/pmoura-dev/gobroker"
	"github.com/pmoura-dev/hauto.device-gateway/controllers"
	"github.com/pmoura-dev/hauto.device-gateway/types"
)

func SetHeaterCoolerMode(ctx gobroker.ConsumerContext, message gobroker.Message) error {
	var action types.SetHeaterCoolerModeMessage
	err := json.Unmarshal(message.GetBody(), &action)
	if err != nil {
		return err
	}

	controller := ctx.Params[controllers.ControllerKey].(controllers.HeaterCooler)

	mode := action.Mode
	data, err := controller.SetHeaterCoolerMode(mode)
	if err != nil {
		return err
	}

	return ctx.Publisher.Publish([]byte(data.Payload), data.Topic, map[string]any{
		"exchange": "amq.topic",
	})
}
