package scryfall

import (
	"cloud.google.com/go/civil"
)

type Preview struct {
	// The name of the source that previewed this card.
	Source string `json:"source"`
	// A link to the preview for this card.
	SourceUri string `json:"source_uri"`
	// The date this card was previewed.
	PreviewedAt civil.Date `json:"previewed_at"`
}
