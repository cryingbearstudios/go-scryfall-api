package scryfall

import (
	"context"

	"resty.dev/v3"
)

const (
	baseURL         = "https://api.scryfall.com"
	userAgentString = "go-scryfall-api/0.1 (Crying Bear Studios)"
)

type ScryfallClient struct {
	*resty.Client
}

func NewClient() *ScryfallClient {
	return &ScryfallClient{
		resty.New().
			SetHeader("User-Agent", userAgentString).
			SetBaseURL(baseURL),
	}
}

func (c *ScryfallClient) Close() error {
	return c.Close()
}

func (c *ScryfallClient) r(ctx context.Context) *resty.Request {
	return c.R().
		SetContext(ctx).
		SetHeader("Accept", "application/json").
		SetError(Error{})
}
