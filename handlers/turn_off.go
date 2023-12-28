package handlers

import (
	"github.com/pmoura-dev/gobroker"
	"github.com/pmoura-dev/hauto.device-gateway/controllers"
)

func TurnOff(ctx gobroker.ConsumerContext, _ gobroker.Message) error {
	controller := ctx.Params[controllers.ControllerKey]
	switchableController := controller.(controllers.Switchable)

	return switchableController.TurnOff()
}
