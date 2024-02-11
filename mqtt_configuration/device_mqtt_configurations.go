package mqtt_configuration

import (
	"github.com/pmoura-dev/hauto.device-gateway/clients/transaction_service"
	"github.com/pmoura-dev/hauto.device-gateway/types"
)

/*
var deviceMQTTConfigurations types.DeviceMQTTConfigurations = map[int]types.DeviceMQTTConfiguration{
	1: {
		Controller: "shelly_bulb",
		Actions: map[string]string{
			types.TurnOnAction:  "shellies.shellycolorbulb-3494546B2CAD.color.0.command",
			types.TurnOffAction: "shellies.shellycolorbulb-3494546B2CAD.color.0.command",
			types.SetRGBColor:   "shellies.shellycolorbulb-3494546B2CAD.color.0.set",
		},
		Listeners: map[string]string{
			types.StateProperty: "shellies.shellycolorbulb-3494546B2CAD.color.0.status",
		},
	},
}*/

var deviceMQTTConfigurations types.DeviceMQTTConfigurations

func GetDeviceMQTTConfigurations() (types.DeviceMQTTConfigurations, error) {
	if deviceMQTTConfigurations == nil {
		var err error
		deviceMQTTConfigurations, err = transaction_service.GetDevicesMQTTConfigurations()
		if err != nil {
			return nil, err
		}
	}

	return deviceMQTTConfigurations, nil
}
