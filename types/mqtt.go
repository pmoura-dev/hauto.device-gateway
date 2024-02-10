package types

type MQTTData struct {
	Topic   string
	Payload string
}

type DeviceMQTTConfigurations map[int]DeviceMQTTConfiguration

type DeviceMQTTConfiguration struct {
	Controller string `json:"controller"`
	Actions    map[string]string
	Listeners  map[string]string
}
