package action_requests

import (
	"encoding/json"

	"github.com/pmoura-dev/gobroker"
	"github.com/pmoura-dev/hauto.device-gateway/api"
)

func PublishActionSucceeded(publisher gobroker.Publisher, metadata api.Metadata, callbackURL string) {
	message := ActionSucceededMessage{
		Meta:        metadata,
		CallbackURL: callbackURL,
	}

	body, _ := json.Marshal(message)

	_ = publisher.Publish(body, api.TopicActionSucceeded, map[string]any{
		"exchange":       api.ExchangeDevices,
		"correlation_id": message.Meta.CorrelationID,
	})
}

func PublishActionFailed(publisher gobroker.Publisher, err error, metadata api.Metadata, callbackURL string) {
	message := ActionFailedMessage{
		Meta:        metadata,
		Error:       err.Error(),
		CallbackURL: callbackURL,
	}

	body, _ := json.Marshal(message)

	_ = publisher.Publish(body, api.TopicActionFailed, map[string]any{
		"exchange":       api.ExchangeDevices,
		"correlation_id": message.Meta.CorrelationID,
	})
}
