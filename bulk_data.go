package scryfall

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"time"

	"github.com/google/uuid"
)

// BulkData represents a Scryfall bulk export file.
//
// Scryfall provides daily exports of our card data in bulk files. Each of these
// files is represented as a bulk_data object via the API. URLs for files change
// their timestamp each day, and can be fetched programmatically.
//
// Please note:
//   - Card objects in bulk data include price information, but prices should be
//     considered dangerously stale after 24 hours. Only use bulk price
//     information to track trends or provide a general estimate of card value.
//     Prices are not updated frequently enough to power a storefront or sales
//     system. You consume price information at your own risk.
//   - Updates to gameplay data (such as card names, Oracle text, mana costs,
//     etc) are much less frequent. If you only need gameplay information,
//     downloading card data once per week or right after set releases would most
//     likely be sufficient.
//   - Every card type in every product is included, including planar cards,
//     schemes, Vanguard cards, tokens, emblems, and funny cards. Make sure
//     you've reviewed documentation for the Card type.
//
// Bulk data is only collected once every 12 hours. You can use the card API
// methods to retrieve fresh objects instead.
type BulkData struct {
	// A content type for this object, always "bulk_data".
	ObjectType string `json:"object"`

	// A unique ID for this bulk item.
	ID uuid.UUID `json:"id"`

	// The Scryfall API URI for this file.
	URI string `json:"uri"`

	// A computer-readable string for the kind of bulk item.
	Type string `json:"type"`

	// A human-readable name for this file.
	Name string `json:"name"`

	// A human-readable description for this file.
	Description string `json:"description"`

	// The URI that hosts this bulk file for fetching.
	DownloadURI string `json:"download_uri"`

	// The time when this file was last updated.
	UpdatedAt time.Time `json:"updated_at"`

	// The size of this file in integer bytes.
	Size int64 `json:"size"`

	// The MIME type of this file.
	ContentType string `json:"content_type"`

	// The Content-Encoding encoding that will be used to transmit this file when you download it.
	ContentEncoding string `json:"content_encoding"`
}

type EnumerationCallback[T any] func(context.Context, *T) error

// LIMITATION: writes progress, but only to the logger, and this is currently not optional.
func (bd BulkData) Enumerate(ctx context.Context, callback EnumerationCallback[Card]) error {
	req, err := http.NewRequestWithContext(ctx, "GET", bd.DownloadURI, nil)
	if err != nil {
		return err
	}
	req.Header["User-Agent"] = []string{userAgentString}
	req.Header["Accept"] = []string{bd.ContentType}
	slog.DebugContext(ctx, "fetch bulk data file", "uri", bd.DownloadURI)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	startingToken, err := decoder.Token()
	if err != nil {
		return err
	}
	if delim, ok := startingToken.(json.Delim); !ok || delim != '[' {
		return fmt.Errorf("expected bulk data file to begin with array start [ but got %v", startingToken)
	}
	slog.DebugContext(ctx, "begin", "fileSize", bd.Size)

	cardCount := 0
	percentComplete := 0
	for decoder.More() {
		var card Card
		if err = decoder.Decode(&card); err != nil {
			return err
		}
		if err = callback(ctx, &card); err != nil {
			return err
		}
		cardCount += 1
		fraction := float32(decoder.InputOffset()) / float32(bd.Size)
		percentage := int(fraction * 100)

		if percentage%5 == 0 && percentage > percentComplete {
			slog.InfoContext(ctx, "progress",
				"count", cardCount,
				"offset", decoder.InputOffset(),
				"percent", percentage)
			percentComplete = percentage
		}
		// check for context cancellation
		if err := ctx.Err(); err != nil {
			return err
		}
	}
	return nil
}

func (c *ScryfallClient) GetBulkDataByType(ctx context.Context, bulkDataType string) (*BulkData, error) {
	var bulkDataEntry BulkData
	urlString, err := url.JoinPath("bulk-data", bulkDataType)
	if err != nil {
		return nil, err
	}
	slog.LogAttrs(ctx, slog.LevelDebug, "request bulk data", slog.String("urlString", urlString))
	resp, err := c.r(ctx).SetResult(&bulkDataEntry).Get(urlString)
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, fmt.Errorf("failed to fetch bulk data for type %s: %v", bulkDataType, resp.Error().(Error).Details)
	}
	return &bulkDataEntry, nil
}
