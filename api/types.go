package api

const (
	ExchangeDevices      = "devices"
	ExchangeDefaultTopic = "amq.topic"
)

const (
	QueueActionRequestedTurnOn  = "action.requested.turn_on.device-gateway.queue"
	QueueActionRequestedTurnOff = "action.requested.turn_off.device-gateway.queue"

	QueueStateUpdatedFormat = "state.updated.%d.device-gateway.queue"
)

const (
	TopicActionRequestedTurnOn  = "action.requested.turn_on"
	TopicActionRequestedTurnOff = "action.requested.turn_off"

	TopicActionSucceeded = "action.succeeded"
	TopicActionFailed    = "action.failed"

	TopicStateUpdated = "state.updated"
)

type Metadata struct {
	CorrelationID string `json:"correlation_id"`
}
