package listeners

import (
	"fmt"
	"log"

	"github.com/pmoura-dev/gobroker"
)

const (
	powerResponseTopicFormat = "power.%s"
	powerResponseExchange    = "devices"
)

func Power(ctx gobroker.ConsumerContext, message gobroker.Message) error {
	naturalID := ctx.Params["natural_id"].(string)

	body := message.GetBody()
	log.Printf("power received from %s: %s", naturalID, string(body))

	responseTopic := fmt.Sprintf(powerResponseTopicFormat, naturalID)
	return ctx.Publisher.Publish(body, responseTopic, map[string]any{
		"exchange": powerResponseExchange,
	})
}
