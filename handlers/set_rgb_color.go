package handlers

import (
	"encoding/json"

	"github.com/pmoura-dev/gobroker"
	"github.com/pmoura-dev/hauto.device-gateway/api"
	"github.com/pmoura-dev/hauto.device-gateway/controllers"
)

func SetRGBColor(ctx gobroker.ConsumerContext, message gobroker.Message) error {
	var action api.SetRGBColorMessage
	err := json.Unmarshal(message.GetBody(), &action)
	if err != nil {
		return err
	}

	controller := ctx.Params[controllers.ControllerKey]
	switchableController := controller.(controllers.RGBColored)

	color := action.Color
	return switchableController.SetRGBColor(color.Red, color.Green, color.Blue)
}
