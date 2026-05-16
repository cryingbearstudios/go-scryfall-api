package scryfall

import (
	"fmt"
	"path"
	"time"
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
	ID string `json:"id"`

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

func (c *ScryfallClient) GetBulkDataByType(bulkDataType string) (*BulkData, error) {
	var bulkDataEntry BulkData
	resp, err := c.r().
		SetResult(&bulkDataEntry).
		Get(path.Join(baseURL, "bulk-data", bulkDataType))
	if err != nil {
		return nil, fmt.Errorf("failed to fetch bulk data for type %s: %w", bulkDataType, err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("failed to fetch bulk data for type %s: %v", bulkDataType, c.errorResponse.Details)
	}
	return &bulkDataEntry, nil
}
