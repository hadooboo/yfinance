package yfinance

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

const (
	yahooFinanceAPIURL = "https://query2.finance.yahoo.com"
	defaultTimeout     = 30 * time.Second
)

type Client struct {
	httpClient *http.Client
}

func New() *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: defaultTimeout,
		},
	}
}

func (r *Client) do(req *http.Request, v interface{}) error {
	resp, err := r.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("received an error response from api: " + resp.Status)
	}

	return json.NewDecoder(resp.Body).Decode(&v)
}
