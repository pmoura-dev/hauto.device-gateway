package types

import (
	"github.com/shopspring/decimal"
)

type LightBulbState struct {
	Status      string          `json:"status"`
	Mode        string          `json:"mode"`
	Hue         decimal.Decimal `json:"hue"`
	Saturation  decimal.Decimal `json:"saturation"`
	Lightness   decimal.Decimal `json:"lightness"`
	Brightness  int             `json:"brightness"`
	Temperature int             `json:"temperature"`
}
