package actions

import (
	"github.com/pmoura-dev/gobroker"
	"github.com/pmoura-dev/hauto.device-gateway/controllers"
)

func TurnOn(ctx gobroker.ConsumerContext, _ gobroker.Message) error {
	controller := ctx.Params[controllers.ControllerKey]
	switchableController := controller.(controllers.Switchable)

	data := switchableController.TurnOn()
	return ctx.Publisher.Publish([]byte(data.Payload), data.Topic, map[string]any{
		"exchange": "amq.topic",
	})
}
