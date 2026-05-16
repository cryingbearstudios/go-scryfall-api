package scryfall

import "time"

type Preview struct {
	// The name of the source that previewed this card.
	Source string `json:"source"`
	// A link to the preview for this card.
	SourceUri string `json:"source_uri"`
	// The date this card was previewed.
	PreviewedAt time.Time `json:"previewed_at"`
}
