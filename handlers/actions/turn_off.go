package actions

import (
	"github.com/pmoura-dev/gobroker"
	"github.com/pmoura-dev/hauto.device-gateway/controllers"
)

func TurnOff(ctx gobroker.ConsumerContext, _ gobroker.Message) error {
	controller := ctx.Params[controllers.ControllerKey].(controllers.Switchable)

	data, err := controller.TurnOff()
	if err != nil {
		return err
	}

	return ctx.Publisher.Publish([]byte(data.Payload), data.Topic, map[string]any{
		"exchange": "amq.topic",
	})
}
