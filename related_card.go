package scryfall

// RelatedCard represents a card that is closely related to another card
// (because it is referenced by name, generates a token, melds, etc.).
type RelatedCard struct {
	// A content type for this object, always "related_card".
	ObjectType string `json:"object"`

	// An unique ID for this card in Scryfall’s database.
	ID string `json:"id"`

	// A field explaining what role this card plays in this relationship.
	Component string `json:"component"`

	// The name of this particular related card.
	Name string `json:"name"`

	// The type line of this card.
	TypeLine string `json:"type_line"`

	// A URI where you can retrieve a full object describing this card on Scryfall’s API.
	URI string `json:"uri"`
}
