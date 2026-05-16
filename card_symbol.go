package scryfall

// CardSymbol represents an illustrated symbol that may appear in a card’s mana
// cost or Oracle text. Symbols are based on the notation used in the
// Comprehensive Rules.
type CardSymbol struct {
	// A content type for this object, always "card_symbol".
	ObjectType string `json:"object"`

	// The plaintext symbol. Often surrounded with curly braces {}. Note that not
	// all symbols are ASCII text (for example, {∞}).
	Symbol string `json:"symbol"`

	// An alternate version of this symbol, if it is possible to write it
	// without curly braces.
	LooseVariant *string `json:"loose_variant,omitempty"`

	// An English snippet that describes this symbol. Appropriate for use in alt
	// text or other accessible communication formats.
	English string `json:"english"`

	// True if it is possible to write this symbol “backwards”. For example, the
	// official symbol {U/P} is sometimes written as {P/U} or {P\U} in informal
	// settings. Note that the Scryfall API never writes symbols backwards in
	// other responses. This field is provided for informational purposes.
	Transposable bool `json:"transposable"`

	// True if this is a mana symbol.
	RepresentsMana bool `json:"represents_mana"`

	// A decimal number representing this symbol’s mana value (also known as the
	// converted mana cost). Note that mana symbols from funny sets can have
	// fractional mana values.
	ManaValue *float64 `json:"mana_value,omitempty"`

	// True if this symbol appears in a mana cost on any Magic card. For example
	// {20} has this field set to false because {20} only appears in Oracle text,
	// not mana costs.
	AppearsInManaCosts bool `json:"appears_in_mana_costs"`

	// True if this symbol is only used on funny cards or Un-cards.
	Funny bool `json:"funny"`

	// An array of colors that this symbol represents.
	Colors []string `json:"colors"`

	// An array of plaintext versions of this symbol that Gatherer uses on old
	// cards to describe original printed text. For example: {W} has ["oW",
	// "ooW"] as alternates.
	GathererAlternates *string `json:"gatherer_alternates,omitempty"`

	// A URI to an SVG image of this symbol on Scryfall’s CDNs.
	SvgURI *string `json:"svg_uri,omitempty"`
}
