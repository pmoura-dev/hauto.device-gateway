package controllers

import (
	"github.com/pmoura-dev/hauto.device-gateway/pkg/states"
)

type Controller interface {
	isController()
}

type BaseController struct {
	naturalID string
}

func (c BaseController) isController() {}

// Listeners

type StateListener interface {
	Controller
	GetStateTopic() string
	HandleState(rawState []byte, topic string) (states.State, error)
}

// Actions

type MQTTData struct {
	Topic   string
	Payload string
}

type Switchable interface {
	Controller
	TurnOn() MQTTData
	TurnOff() MQTTData
}
