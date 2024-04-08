package handlers

import (
	"encoding/json"

	"github.com/pmoura-dev/gobroker"
	"github.com/pmoura-dev/hauto.device-gateway/api"
	action_requests2 "github.com/pmoura-dev/hauto.device-gateway/api/action_requests"
	"github.com/pmoura-dev/hauto.device-gateway/pkg/controllers"
	"github.com/pmoura-dev/hauto.device-gateway/pkg/memory"
)

type TurnOffMessage struct {
	action_requests2.ActionRequestedMessage
}

func TurnOff(ctx gobroker.ConsumerContext, message gobroker.Message) error {
	var action TurnOnMessage
	if err := json.Unmarshal(message.GetBody(), &action); err != nil {
		return err
	}

	manager, err := memory.GetDeviceManagerInstance()
	if err != nil {
		return err
	}

	baseController := manager.GetController(action.DeviceID)
	if baseController == nil {
		err = action_requests2.ErrControllerNotImplemented
		action_requests2.PublishActionFailed(ctx.Publisher, err, action.Meta, action.CallbackURL)
		return err
	}

	controller, ok := baseController.(controllers.Switchable)
	if !ok {
		err = action_requests2.ErrControllerNotImplemented
		action_requests2.PublishActionFailed(ctx.Publisher, err, action.Meta, action.CallbackURL)
		return err
	}

	mqttData := controller.TurnOff()
	err = ctx.Publisher.Publish([]byte(mqttData.Payload), mqttData.Topic, map[string]any{
		"correlation_id": action.Meta.CorrelationID,
		"exchange":       api.ExchangeDefaultTopic,
	})
	if err != nil {
		action_requests2.PublishActionFailed(ctx.Publisher, err, action.Meta, action.CallbackURL)
		return err
	}

	action_requests2.PublishActionSucceeded(ctx.Publisher, action.Meta, action.CallbackURL)
	return nil
}
