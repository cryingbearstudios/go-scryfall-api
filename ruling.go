package scryfall

import "time"

type Ruling struct {
	// A content type for this object, always "ruling".
	ObjectType string `json:"object"`

	// A computer-readable string indicating which company produced this ruling.
	Source string `json:"source"`

	/// The date when the ruling or note was published.
	PublishedAt time.Time `json:"published_at,omitempty"`

	/// The text of the ruling.
	Comment string `json:"comment"`
}
