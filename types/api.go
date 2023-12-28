package types

type BaseActionMessage struct {
	NaturalID  string `json:"natural_id"`
	Controller string `json:"controller"`
}

type SetRGBColorMessage struct {
	BaseActionMessage
	Color struct {
		Red   int
		Green int
		Blue  int
	}
}
