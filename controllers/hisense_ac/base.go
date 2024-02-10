package hisense_ac

type HisenseACController struct {
	actions   map[string]string
	listeners map[string]string
}

func NewHisenseACController(actions, listeners map[string]string) HisenseACController {
	return HisenseACController{
		actions:   actions,
		listeners: listeners,
	}
}
