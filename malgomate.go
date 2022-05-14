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

// Client is the main malgomate wrapper. It holds the HTTP client, the MAL API URL, as well as your
// MAL API key. Recommended to intialize via the NewClient constructor, but you can choose to construct
// by hand incase you need to do some overrides/injection.
type Client struct {
	BaseURL    string
	apiKey     string
	HTTPClient *http.Client
}

// NewClient is a constructor for quickly building the malgomate client. Requires you to pass your
// API key value and will generate with some sane default values.
func NewClient(apiKey string) *Client {
	return &Client{
		apiKey:  apiKey,
		BaseURL: BaseURLv2,
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// Paging will be present when a response has additional result items
type Paging struct {
	Next     string `json:"next"`
	Previous string `json:"previous"`
}

// HasNext checks to see if there is a next page available
func (p *Paging) HasNext() bool {
	return len(p.Next) != 0
}

// HasPrevious checks to see if there is a previous page available
func (p *Paging) HasPrevious() bool {
	return len(p.Previous) != 0
}

// GetNextPage is a helper function that will automatically retrieve the next page
// of data, if one is present.
func (c *Client) GetNextPage(p *Paging, v interface{}) error {
	if !p.HasNext() {
		return errors.New("no next page to fetch")
	}
	req, err := http.NewRequest(http.MethodGet, p.Next, nil)
	if err != nil {
		return err
	}

	return c.sendRequest(req, &v)
}

// errorResponse is a general eror wrapper
type errorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

// sendRequest handles all outgoing requests. Takes in an HTTP request and a reference
// to the resulting object. sendRequest will make the API call, handle any error responses,
// and decode the response message into the specified value
func (c *Client) sendRequest(req *http.Request, value interface{}) error {
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

	if err = json.NewDecoder(res.Body).Decode(&value); err != nil {
		return err
	}

	return nil
}
