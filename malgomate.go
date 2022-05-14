// Package malgomate provies an interface to the public MAL (MyAnimeList) API.
package malgomate

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

const (
	BaseURLv2 = "https://api.myanimelist.net/v2"
)

type Client struct {
	BaseURL    string
	apiKey     string
	HTTPClient *http.Client
}

func NewClient(apiKey string) *Client {
	return &Client{
		apiKey:  apiKey,
		BaseURL: BaseURLv2,
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

type Paging struct {
	Next     string `json:"next"`
	Previous string `json:"previous"`
}

func (p *Paging) HasNext() bool {
	return len(p.Next) != 0
}

func (c *Client) GetNextPage(p *Paging, v interface{}) error {
	req, err := http.NewRequest(http.MethodGet, p.Next, nil)
	if err != nil {
		return err
	}

	return c.sendRequest(req, &v)
}

type errorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Dummy   string `json:"bug,omitempty"`
}

func (c *Client) sendRequest(req *http.Request, v interface{}) error {
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json; charset=utf-8")
	req.Header.Set("X-MAL-CLIENT-ID", c.apiKey)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode >= 400 {
		var errRes errorResponse
		if err = json.NewDecoder(res.Body).Decode(&errRes); err == nil {
			return errors.New(errRes.Message)
		}
		return fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}

	if err = json.NewDecoder(res.Body).Decode(&v); err != nil {
		return err
	}

	return nil
}
