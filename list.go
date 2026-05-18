package scryfall

import (
	"context"
	"fmt"
)

// List represents a requested sequence of other objects (Cards, Sets, etc).
// List objects may be paginated, and also include information about issues raised
// when generating the list.
type List[T any] struct {
	// A content type for this object, always "list".
	ObjectType string `json:"object"`

	// An array of the requested objects, in a specific order.
	Data []T `json:"data"`

	// True if this List is paginated and there is a page beyond the current page.
	HasMore bool `json:"has_more"`

	// If there is a page beyond the current page, this field will contain a full
	// API URI to that page. You may submit a HTTP GET request to that URI to
	// continue paginating forward on this List.
	NextPage string `json:"next_page,omitempty"`

	// If this is a list of Card objects, this field will contain the total number
	// of cards found across all pages.
	TotalCards *int `json:"total_cards,omitempty"`

	// An array of human-readable warnings issued when generating this list, as
	// strings. Warnings are non-fatal issues that the API discovered with your
	// input. In general, they indicate that the List will not contain all of the
	// information you requested. You should fix the warnings and re-submit your
	// request.
	Warnings []string `json:"warnings,omitempty"`
}

type PaginationCallback[T any] func(context.Context, T) error

func (l *List[T]) Paginate(ctx context.Context, client *ScryfallClient, callback PaginationCallback[T]) error {
	for _, item := range l.Data {
		err := callback(ctx, item)
		if err != nil {
			return err
		}
	}
	if !l.HasMore {
		return nil
	}
	nextPage := List[T]{}
	resp, err := client.r(ctx).SetResult(nextPage).Get(l.NextPage)
	if err != nil {
		return err
	}
	if resp.IsError() {
		return fmt.Errorf("failed to fetch during pagination on url %s: %v", l.NextPage, resp.Error().(Error).Details)
	}
	return nextPage.Paginate(ctx, client, callback)
}
