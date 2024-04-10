package transaction_service

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

var baseURL string

func Setup(url string) {
	baseURL = url
}

func execute(procName string, params map[string]any) ([]byte, error) {

	data, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	response, err := http.Post(baseURL+"/execute/"+procName, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	_ = response.Body.Close()

	return body, nil
}
