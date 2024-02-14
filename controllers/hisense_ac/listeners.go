package hisense_ac

import (
	"encoding/json"

	"github.com/pmoura-dev/hauto.device-gateway/types"
)

func (c HisenseACController) HandleState(rawState []byte) ([]byte, error) {
	var state hisenseACState

	err := json.Unmarshal(rawState, &state)
	if err != nil {
		return nil, err
	}

	status := "offline"
	if state.TPower == "ON" {
		status = "online"
	}

	mode := "automatic"
	switch state.TWorkMode {
	case "HEAT":
		mode = "heating"
	case "COOL":
		mode = "cooling"
	}

	convertedState := types.AirConditionerState{
		Status:                             status,
		Mode:                               mode,
		CurrentTemperature:                 state.FTempIn,
		CurrentHeatingThresholdTemperature: state.TTemp,
		CurrentCoolingThresholdTemperature: state.TTemp,
	}

	return json.Marshal(convertedState)
}

type hisenseACState struct {
	FTempIn   int    `json:"f_temp_in"`
	TPower    string `json:"t_power"`
	TWorkMode string `json:"t_work_mode"`
	TTemp     int    `json:"t_temp"`
}
