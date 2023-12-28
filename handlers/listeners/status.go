package listeners

import (
	"fmt"
	"log"

	"github.com/pmoura-dev/gobroker"
)

const (
	statusResponseTopicFormat = "status.%s"
	statusResponseExchange    = "devices"
)

func Status(ctx gobroker.ConsumerContext, message gobroker.Message) error {
	naturalID := ctx.Params["natural_id"].(string)

	body := message.GetBody()
	log.Printf("status received from %s: %s", naturalID, string(body))

	responseTopic := fmt.Sprintf(statusResponseTopicFormat, naturalID)
	return ctx.Publisher.Publish(body, responseTopic, map[string]any{
		"exchange": statusResponseExchange,
	})
}
