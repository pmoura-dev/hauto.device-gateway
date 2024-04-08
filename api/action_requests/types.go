package action_requests

import (
	"errors"

	"github.com/pmoura-dev/hauto.device-gateway/api"
)

type ActionRequestedMessage struct {
	Meta        api.Metadata `json:"meta"`
	DeviceID    int          `json:"device_id"`
	CallbackURL string       `json:"callback_url"`
}

type ActionSucceededMessage struct {
	Meta        api.Metadata `json:"meta"`
	CallbackURL string       `json:"callback_url,omitempty"`
}

type ActionFailedMessage struct {
	Meta        api.Metadata `json:"meta"`
	Error       string       `json:"error"`
	CallbackURL string       `json:"callback_url,omitempty"`
}

var (
	ErrControllerNotImplemented = errors.New("controller is not implemented")
)
