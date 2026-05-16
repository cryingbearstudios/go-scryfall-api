package scryfall

import (
	"context"

	"resty.dev/v3"
)

const (
	// baseURL is the base URL for all Scryfall API endpoints.
	baseURL = "https://api.scryfall.com"
)

type ScryfallClient struct {
	rc            *resty.Client
	errorResponse Error
}

func NewClient(ctx context.Context) *ScryfallClient {
	return &ScryfallClient{
		rc: resty.New().SetContext(ctx),
	}
}

func (c *ScryfallClient) Close() {
	c.rc.Close()
}

func (c *ScryfallClient) r() *resty.Request {
	return c.rc.R().
		SetHeader("Accept", "application/json").
		SetError(c.errorResponse)
}
