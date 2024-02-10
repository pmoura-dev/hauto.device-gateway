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
