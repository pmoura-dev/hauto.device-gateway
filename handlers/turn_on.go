package handlers

import (
	"fmt"

	"github.com/pmoura-dev/gobroker"
	"github.com/pmoura-dev/hauto.device-gateway/controllers"
)

func TurnOn(ctx gobroker.ConsumerContext, _ gobroker.Message) error {
	fmt.Println("yoooo")
	controller := ctx.Params[controllers.ControllerKey]
	switchableController := controller.(controllers.Switchable)

	return switchableController.TurnOn()
}
