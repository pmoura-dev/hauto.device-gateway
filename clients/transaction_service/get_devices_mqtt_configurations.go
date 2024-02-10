package transaction_service

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/pmoura-dev/hauto.device-gateway/types"
)

func GetDevicesMQTTConfigurations() (types.DeviceMQTTConfigurations, error) {
	path := "/get_devices_listeners_with_controllers"

	response, err := http.Get(baseURL + path)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var configurations types.DeviceMQTTConfigurations
	err = json.Unmarshal(body, &configurations)
	if err != nil {
		return nil, err
	}

	return configurations, nil
}
