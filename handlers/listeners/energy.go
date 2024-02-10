package listeners

import (
	"fmt"
	"log"

	"github.com/pmoura-dev/gobroker"
)

const (
	energyResponseTopicFormat = "energy.%s"
	energyResponseExchange    = "devices"
)

func Energy(ctx gobroker.ConsumerContext, message gobroker.Message) error {
	naturalID := ctx.Params["natural_id"].(string)

	body := message.GetBody()
	log.Printf("energy received from %s: %s", naturalID, string(body))

	responseTopic := fmt.Sprintf(energyResponseTopicFormat, naturalID)
	return ctx.Publisher.Publish(body, responseTopic, map[string]any{
		"exchange": energyResponseExchange,
	})
}
