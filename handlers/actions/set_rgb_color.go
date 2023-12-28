package actions

import (
	"encoding/json"

	"github.com/pmoura-dev/gobroker"
	"github.com/pmoura-dev/hauto.device-gateway/controllers"
	"github.com/pmoura-dev/hauto.device-gateway/types"
)

func SetRGBColor(ctx gobroker.ConsumerContext, message gobroker.Message) error {
	var action types.SetRGBColorMessage
	err := json.Unmarshal(message.GetBody(), &action)
	if err != nil {
		return err
	}

	controller := ctx.Params[controllers.ControllerKey]
	rgbColoredController := controller.(controllers.RGBColored)

	color := action.Color
	data := rgbColoredController.SetRGBColor(color.Red, color.Green, color.Blue)
	return ctx.Publisher.Publish([]byte(data.Payload), data.Topic, map[string]any{
		"exchange": "amq.topic",
	})
}
