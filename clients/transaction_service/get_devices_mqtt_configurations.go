package transaction_service

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/pmoura-dev/hauto.device-gateway/types"
)

type deviceMQTTConfigurations map[string]types.DeviceMQTTConfiguration

func GetDevicesMQTTConfigurations() (types.DeviceMQTTConfigurations, error) {
	path := "/execute/get_devices_mqtt_configuration"

	response, err := http.Post(baseURL+path, "", nil)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var tempConfigurations deviceMQTTConfigurations
	err = json.Unmarshal(body, &tempConfigurations)
	if err != nil {
		return nil, err
	}

	configurations := make(types.DeviceMQTTConfigurations)
	for k, v := range tempConfigurations {
		deviceID, err := strconv.Atoi(k)
		if err != nil {
			return nil, err
		}
		configurations[deviceID] = v
	}

	return configurations, nil
}
