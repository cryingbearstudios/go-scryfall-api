package scryfall

import "cloud.google.com/go/civil"

type Ruling struct {
	// A content type for this object, always "ruling".
	ObjectType string `json:"object"`

	// A computer-readable string indicating which company produced this ruling.
	Source string `json:"source"`

	/// The date when the ruling or note was published.
	PublishedAt civil.Date `json:"published_at,omitempty"`

	/// The text of the ruling.
	Comment string `json:"comment"`
}
