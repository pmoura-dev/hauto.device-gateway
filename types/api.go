package types

type BaseActionMessage struct {
	DeviceID int `json:"device_id"`
}

type SetRGBColorMessage struct {
	BaseActionMessage
	Color struct {
		Red   int
		Green int
		Blue  int
	}
}

type SetTemperatureThresholdMessage struct {
	BaseActionMessage
	Value int `json:"value"`
}

type SetHeaterCoolerModeMessage struct {
	BaseActionMessage
	Mode string `json:"mode"`
}
