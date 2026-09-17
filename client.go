package scryfall

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
)

const (
	BaseURL         = "https://api.scryfall.com"
	UserAgentString = "go-scryfall-api/0.1 (Crying Bear Studios)"
)

type ScryfallClient struct {
	c *http.Client
}

func NewClient() *ScryfallClient {
	return &ScryfallClient{
		&http.Client{},
	}
}

func (c *ScryfallClient) r(ctx context.Context) *http.Request {
	url, err := url.Parse(BaseURL)
	if err != nil {
		panic(err)
	}
	req := http.Request{
		URL:        url,
		Proto:      "",
		ProtoMajor: 0,
		ProtoMinor: 0,
		Header: http.Header{
			"Accept":     []string{"application/json"},
			"User-Agent": []string{UserAgentString},
		},
	}
	return req.WithContext(ctx)
}

func get[successType any](ctx context.Context, c *ScryfallClient, elem ...string) (*successType, *Error, error) {
	url, err := url.JoinPath(BaseURL, elem...)
	if err != nil {
		return nil, nil, err
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", UserAgentString)
	req = req.WithContext(ctx)
	resp, err := c.c.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	var successPayload successType
	var errorPayload Error
	var val any = &successPayload
	if resp.StatusCode >= 400 {
		val = &errorPayload
	}
	if err := decoder.Decode(val); err != nil {
		return nil, nil, err
	}
	if resp.StatusCode >= 400 {
		return nil, &errorPayload, nil
	}
	return &successPayload, nil, nil
}
