package scryfall

import (
	"context"

	"resty.dev/v3"
)

const (
	BaseURL         = "https://api.scryfall.com"
	UserAgentString = "go-scryfall-api/0.1 (Crying Bear Studios)"
)

type ScryfallClient struct {
	c *resty.Client
}

func NewClient() *ScryfallClient {
	return &ScryfallClient{
		resty.New().
			SetHeader("User-Agent", UserAgentString).
			SetBaseURL(BaseURL),
	}
}

func (c *ScryfallClient) r(ctx context.Context) *resty.Request {
	return c.c.R().
		SetContext(ctx).
		SetHeader("Accept", "application/json").
		SetError(Error{})
}
