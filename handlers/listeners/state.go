package listeners

import (
	"fmt"
	"log"

	"github.com/pmoura-dev/gobroker"
	"github.com/pmoura-dev/hauto.device-gateway/controllers"
)

const (
	stateResponseTopicFormat = "state.%d"
	stateResponseExchange    = "devices"
)

func State(ctx gobroker.ConsumerContext, message gobroker.Message) error {

	deviceID := ctx.Params["device_id"].(int)
	controller := ctx.Params[controllers.ControllerKey].(controllers.StateListener)

	body := message.GetBody()
	log.Printf("raw state received from device %d: %s", deviceID, string(body))

	state, err := controller.HandleState(body)
	if err != nil {
		return err
	}

	responseTopic := fmt.Sprintf(stateResponseTopicFormat, deviceID)
	return ctx.Publisher.Publish(state, responseTopic, map[string]any{
		"exchange": stateResponseExchange,
	})
}
