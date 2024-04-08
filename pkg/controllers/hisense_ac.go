package controllers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pmoura-dev/hauto.device-gateway/pkg/states"
	"github.com/pmoura-dev/hauto.device-gateway/pkg/types"
)

const (
	HisenseACControllerName = "hisense_ac"
)

type HisenseACController struct {
	BaseController
}

func NewHisenseACController(naturalID string) HisenseACController {
	return HisenseACController{
		BaseController{naturalID: naturalID},
	}
}

func (c HisenseACController) GetStateTopic() string {
	const topicFormat = "hisense_ac.%s.*.status"

	return fmt.Sprintf(topicFormat, c.naturalID)
}

func (c HisenseACController) HandleState(rawState []byte, topic string) (states.State, error) {
	value := string(rawState)

	switch {
	case strings.Contains(topic, "t_power"):
		status := types.OfflineStatus
		if value == "ON" {
			status = types.OnlineStatus
		}

		return states.SingleParamState{
			Name:  "status",
			Value: status,
		}, nil

	case strings.Contains(topic, "t_work_mode"):
		mode := types.AirConditionerAutomaticMode
		switch value {
		case "HEAT":
			mode = types.AirConditionerHeatingMode
		case "COOL":
			mode = types.AirConditionerCoolingMode
		}

		return states.SingleParamState{
			Name:  "mode",
			Value: mode,
		}, nil

	case strings.Contains(topic, "f_temp_in"):
		currentTemperature, err := strconv.Atoi(value)
		if err != nil {
			return nil, err
		}

		return states.SingleParamState{
			Name:  "current_temperature",
			Value: currentTemperature,
		}, nil

	case strings.Contains(topic, "t_temp"):
		currentThresholdTemperature, err := strconv.Atoi(value)
		if err != nil {
			return nil, err
		}

		return states.SingleParamState{
			Name:  "current_threshold_temperature",
			Value: currentThresholdTemperature,
		}, nil
	}

	return states.SingleParamState{}, nil
}
