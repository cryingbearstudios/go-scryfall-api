package scryfall

// ManaCost represents a parsed mana cost from the Scryfall API.
type ManaCost struct {
	// A content type for this object, always "mana_cost".
	ObjectType string `json:"object"`

	// The mana string to parse.
	Cost string `json:"cost"`

	// The mana value. If you submit Un-set mana symbols, this decimal could
	// include fractional parts.
	Cmc float64 `json:"cmc"`

	// The colors of the given cost.
	Colors []string `json:"colors"`

	// True if the cost is colorless.
	Colorless bool `json:"colorless"`

	// True if the cost is monocolored.
	Monocolored bool `json:"monocolored"`

	// True if the cost is multicolored.
	Multicolored bool `json:"multicolored"`
}
