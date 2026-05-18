package scryfall

// Multiface cards have a CardFaces property containing at least two Card Face objects.
type CardFace struct {
	// A content type for this object, always "card_face".
	ObjectType string `json:"object"`

	// The name of the illustrator of this card face. Newly spoiled cards may not have this field yet.
	Artist string `json:"artist,omitempty"`

	// The mana value of this particular face, if the card is reversible.
	Cmc *float64 `json:"cmc,omitempty"`

	// The colors in this face’s color indicator, if any.
	ColorIndicator []string `json:"color_indicator,omitempty"`

	// This face’s colors, if the game defines colors for the individual face of this card.
	Colors []string `json:"colors,omitempty"`

	// The flavor text printed on this face, if any.
	FlavorText string `json:"flavor_text,omitempty"`

	// A unique identifier for the card face artwork that remains consistent across reprints.
	// Newly spoiled cards may not have this field yet.
	IllustrationID string `json:"illustration_id,omitempty"`

	// An object providing URIs to imagery for this face, if this is a double-sided card.
	// If this card is not double-sided, then the image_uris property will be part of the parent object instead.
	ImageURIs map[string]string `json:"image_uris,omitempty"`

	// The layout of this card face, if the card is reversible.
	Layout string `json:"layout,omitempty"`

	// This face’s loyalty, if any.
	Loyalty string `json:"loyalty,omitempty"`

	// The mana cost for this face. This value will be any empty string "" if the cost is absent.
	// Remember that per the game rules, a missing mana cost and a mana cost of {0} are different values.
	ManaCost string `json:"mana_cost"`

	// The name of this particular face.
	Name string `json:"name"`

	// The Oracle ID of this particular face, if the card is reversible.
	OracleID string `json:"oracle_id,omitempty"`

	// The Oracle text for this face, if any.
	OracleText string `json:"oracle_text,omitempty"`

	// This face’s power, if any. Note that some cards have powers that are not numeric, such as *.
	Power string `json:"power,omitempty"`

	// The localized name printed on this face, if any.
	PrintedName string `json:"printed_name,omitempty"`

	// The localized text printed on this face, if any.
	PrintedText string `json:"printed_text,omitempty"`

	// The localized type line printed on this face, if any.
	PrintedTypeLine string `json:"printed_type_line,omitempty"`

	// This face’s toughness, if any.
	Toughness string `json:"toughness,omitempty"`

	// The type line of this particular face, if the card is reversible.
	TypeLine string `json:"type_line,omitempty"`

	// The watermark on this particular card face, if any.
	Watermark string `json:"watermark,omitempty"`

	// Undocumented by Scryfall fields
	ArtistID string `json:"artist_id,omitempty"`
	Defense  string `json:"defense,omitempty"`

	// The just-for-fun name printed on the card face (such as for Godzilla series cards).
	FlavorName string `json:"flavor_name,omitempty"`
}
