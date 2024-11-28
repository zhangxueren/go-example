package base

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type WxExecutorActionResponse struct {
	ErrNo  int                    `json:"errNo"`
	ErrStr string                 `json:"errStr"`
	Data   map[string]interface{} `json:"data"`
}

// RunAction sends an action to the specified URL and returns the response data.
func RunAction(wxId, bizRequestId string, actionType int, actionContent map[string]interface{}) (map[string]interface{}, error) {
	url := "http://10.110.171.84:8099/wxexecutor/action/send"

	// Create the JSON payload.
	payload := map[string]interface{}{
		"wxId":          wxId,
		"bizRequestId":  bizRequestId,
		"bizLine":       "laxinqk", // This is hardcoded here but can be parameterized if needed.
		"actionType":    actionType,
		"actionContent": actionContent,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	// Make the HTTP request.
	req, err := http.NewRequestWithContext(context.Background(), "POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Parse the JSON response.
	var response WxExecutorActionResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	if response.ErrNo != 0 {
		return nil, fmt.Errorf("RunAction error, errNo[%d] errStr[%s]", response.ErrNo, response.ErrStr)
	}

	return response.Data, nil
}
