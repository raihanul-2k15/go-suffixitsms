package suffixitsms

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"
)

type client struct {
	apiKey  string
	baseUrl string

	timeout    time.Duration
	httpClient *http.Client
}

func NewClient(apiKey string) *client {
	return &client{
		apiKey:     apiKey,
		baseUrl:    "https://bulkmsg.suffixit.com:7070/api",
		httpClient: &http.Client{Timeout: 30 * time.Second},
	}
}

func (c *client) SetTimeout(dur time.Duration) *client {
	c.timeout = dur
	c.httpClient.Timeout = dur
	return c
}

func (c *client) sendRequest(method string, path string, payload map[string]string) (string, error) {
	if c.apiKey == "" {
		return "", errors.New("apiKey is empty")
	}

	plBytes, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(method, c.baseUrl+path, bytes.NewBuffer(plBytes))
	if err != nil {
		return "", c.safeError(err)
	}

	req.Header.Add("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", c.safeError(err)
	}

	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", c.safeError(err)
	}

	body := string(bytes)

	if resp.StatusCode < 200 || resp.StatusCode > 202 {
		return "", c.safeError(extractError(body))
	}

	return body, nil
}
