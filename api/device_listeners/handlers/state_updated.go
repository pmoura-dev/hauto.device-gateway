package handlers

import (
	"encoding/json"
	"time"

	"github.com/pmoura-dev/gobroker"
	"github.com/pmoura-dev/hauto.device-gateway/api"
	"github.com/pmoura-dev/hauto.device-gateway/pkg/controllers"
	"github.com/pmoura-dev/hauto.device-gateway/pkg/memory"
)

type StateUpdatedMessage struct {
	DeviceID  int       `json:"device_id"`
	Timestamp time.Time `json:"timestamp"`
	State     any       `json:"state"`
}

func StateUpdated(ctx gobroker.ConsumerContext, message gobroker.Message) error {
	deviceID := ctx.Params["device_id"].(int)
	controller := ctx.Params["controller"].(controllers.StateListener)

	body := message.GetBody()
	topic := message.GetTopic()

	state, err := controller.HandleState(body, topic)
	if err != nil {
		return err
	}

	deviceManager, err := memory.GetDeviceManagerInstance()
	if err != nil {
		return err
	}

	ok, updatedState, err := deviceManager.UpdateState(deviceID, state)
	if err != nil {
		return err
	}
	if !ok {
		return nil
	}

	stateUpdatedMessage := StateUpdatedMessage{
		DeviceID:  deviceID,
		Timestamp: time.Now(),
		State:     updatedState,
	}

	payload, err := json.Marshal(stateUpdatedMessage)
	if err != nil {
		return err
	}

	return ctx.Publisher.Publish(payload, api.TopicStateUpdated, map[string]any{
		"exchange": api.ExchangeDevices,
	})
}
