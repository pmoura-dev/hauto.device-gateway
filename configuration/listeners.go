package configuration

import (
	"os"

	"gopkg.in/yaml.v3"
)

const (
	listenersFileName = "listeners.yaml"
)

type Listeners struct {
	Devices []DeviceListener `yaml:"devices"`
}

type DeviceListener struct {
	NaturalID string  `yaml:"id"`
	Topics    []Topic `yaml:"topics"`
}

type Topic struct {
	Type string `yaml:"type"`
	Name string `yaml:"name"`
}

func LoadDeviceListeners() (Listeners, error) {
	var listeners Listeners

	content, err := os.ReadFile(listenersFileName)
	if err != nil {
		return Listeners{}, err
	}

	err = yaml.Unmarshal(content, &listeners)
	if err != nil {
		return Listeners{}, err
	}

	return listeners, nil
}
