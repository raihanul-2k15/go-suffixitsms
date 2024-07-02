package suffixitsms

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type errorResponse struct {
	Title string `json:"title"`
}

func extractError(resp string) error {
	var jsonResp errorResponse
	err := json.Unmarshal([]byte(resp), &jsonResp)
	if err != nil || jsonResp.Title == "" {
		return fmt.Errorf("API Error: %s, Original Response: %s", "Unknown error", resp)
	}

	return fmt.Errorf("API Error: %s, Original Response: %s", jsonResp.Title, resp)
}

func (c *client) safeError(err error) error {
	if err == nil {
		return nil
	}

	errMsg := strings.Replace(err.Error(), c.apiKey, "********", -1)
	return errors.New(errMsg)
}
