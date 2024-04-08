package transaction_service

import (
	"encoding/json"

	"github.com/pmoura-dev/hauto.device-gateway/data/stubs"
	"github.com/pmoura-dev/hauto.device-gateway/pkg/states"
)

func GetDevicesWithState() ([]DeviceWithState, error) {

	// TODO: make call to transaction service
	body := stubs.DevicesData

	var response []DeviceWithState
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return response, nil
}

var stateTypeMap = map[string]func() states.State{
	"color_light":     func() states.State { return &states.ColorLightState{} },
	"air_conditioner": func() states.State { return &states.AirConditionerState{} },
}

type DeviceWithState struct {
	ID         int    `json:"id"`
	NaturalID  string `json:"natural_id"`
	Type       string `json:"type"`
	Controller string `json:"controller"`
	State      states.State
}

func (d *DeviceWithState) UnmarshalJSON(data []byte) error {
	var device struct {
		ID         int    `json:"id"`
		NaturalID  string `json:"natural_id"`
		Type       string `json:"type"`
		Controller string `json:"controller"`
	}

	if err := json.Unmarshal(data, &device); err != nil {
		return err
	}

	d.ID = device.ID
	d.NaturalID = device.NaturalID
	d.Type = device.Type
	d.Controller = device.Controller

	type rawStateType struct {
		State json.RawMessage `json:"state"`
	}

	var rawState rawStateType
	if err := json.Unmarshal(data, &rawState); err != nil {
		return err
	}

	stateFn, exists := stateTypeMap[d.Type]
	if !exists {
		return nil
	}

	state := stateFn()
	if err := json.Unmarshal(rawState.State, state); err != nil {
		return err
	}

	d.State = state

	return nil
}
