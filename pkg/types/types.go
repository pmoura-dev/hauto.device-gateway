package types

const (
	OnlineStatus  = "online"
	OfflineStatus = "offline"
)

const (
	AirConditionerAutomaticMode = "automatic"
	AirConditionerHeatingMode   = "heating"
	AirConditionerCoolingMode   = "cooling"
)

const (
	ColorLightColorMode = "color"
	ColorLightWhiteMode = "white"
)

type HSLAColor struct {
	Hue        int      `json:"hue"`
	Saturation int      `json:"saturation"`
	Lightness  int      `json:"lightness"`
	Alpha      *float32 `json:"alpha,omitempty"`
}
