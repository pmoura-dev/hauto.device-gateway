package states

import (
	"errors"

	"github.com/pmoura-dev/hauto.device-gateway/pkg/types"
)

type State interface {
	isState()

	UpdateSingleValue(field string, value any) (bool, error)
}

type SingleParamState struct {
	Name  string
	Value any
}

func (s SingleParamState) isState() {}

func (s SingleParamState) UpdateSingleValue(field string, value any) (bool, error) {
	panic("not implemented")
}

type ColorLightState struct {
	Status      string          `json:"status"`
	Mode        string          `json:"mode"`
	Color       types.HSLAColor `json:"color"`
	Brightness  int             `json:"brightness"`
	Temperature int             `json:"temperature"`
}

func (s *ColorLightState) isState() {}

func (s *ColorLightState) UpdateSingleValue(field string, value any) (_ bool, err error) {

	defer func() {
		if p := recover(); p != nil {
			err = errors.New("invalid value")
		}
	}()

	switch field {
	case "status":
		status := value.(string)
		if status == s.Status {
			return false, nil
		}

		s.Status = status
	case "mode":
		mode := value.(string)
		if mode == s.Mode {
			return false, nil
		}

		s.Mode = mode
	case "color":
		color := value.(types.HSLAColor)
		if color == s.Color {
			return false, nil
		}

		s.Color = color
	case "brightness":
		brightness := value.(int)
		if brightness == s.Brightness {
			return false, nil
		}

		s.Brightness = brightness
	case "temperature":
		temperature := value.(int)
		if temperature == s.Temperature {
			return false, nil
		}

		s.Temperature = temperature
	default:
		return false, nil
	}

	return true, nil
}

type AirConditionerState struct {
	Status                      string `json:"status"`
	Mode                        string `json:"mode"`
	CurrentTemperature          int    `json:"current_temperature"`
	CurrentThresholdTemperature int    `json:"current_threshold_temperature"`
}

func (s *AirConditionerState) isState() {}

func (s *AirConditionerState) UpdateSingleValue(field string, value any) (_ bool, err error) {

	defer func() {
		if p := recover(); p != nil {
			err = errors.New("invalid value")
		}
	}()

	switch field {
	case "status":
		status := value.(string)
		if status == s.Status {
			return false, nil
		}

		s.Status = status
	case "mode":
		mode := value.(string)
		if mode == s.Mode {
			return false, nil
		}

		s.Mode = mode
	case "current_temperature":
		currentTemperature := value.(int)
		if currentTemperature == s.CurrentTemperature {
			return false, nil
		}

		s.CurrentTemperature = currentTemperature
	case "current_threshold_temperature":
		currentTemperatureThreshold := value.(int)
		if currentTemperatureThreshold == s.CurrentThresholdTemperature {
			return false, nil
		}

		s.CurrentThresholdTemperature = currentTemperatureThreshold
	default:
		return false, nil
	}

	return true, nil
}
