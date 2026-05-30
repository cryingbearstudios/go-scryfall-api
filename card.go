package scryfall

import (
	"context"
	"fmt"
	"log/slog"
	"net/url"

	"cloud.google.com/go/civil"
	"github.com/google/uuid"
)

// Card objects represent individual Magic: The Gathering cards that players could obtain
// and add to their collection (with a few minor exceptions).
type Card struct {
	// A content type for this object, always "card".
	ObjectType string `json:"object"`

	// This card’s Arena ID, if any. A large percentage of cards are not available on Arena and do not have this ID.
	ArenaID *int64 `json:"arena_id,omitempty"`

	// A unique ID for this card in Scryfall’s database.
	ID uuid.UUID `json:"id"`

	// A language code for this printing.
	Lang string `json:"lang"`

	// This card’s Magic Online ID (also known as the Catalog ID), if any. A large percentage of cards are not available on Magic Online and do not have this ID.
	MtgoID *int64 `json:"mtgo_id,omitempty"`

	// This card’s foil Magic Online ID (also known as the Catalog ID), if any. A large percentage of cards are not available on Magic Online and do not have this ID.
	MtgoFoilID *int64 `json:"mtgo_foil_id,omitempty"`

	// This card’s multiverse IDs on Gatherer, if any, as an array of integers.
	MultiverseIDs []int64 `json:"multiverse_ids,omitempty"`

	// This card’s ID on TCGplayer’s API, also known as the productId.
	TcgplayerID *int64 `json:"tcgplayer_id,omitempty"`

	// This card’s ID on TCGplayer’s API, for its etched version if that version is a separate product.
	TcgplayerEtchedID *int64 `json:"tcgplayer_etched_id,omitempty"`

	// This card’s ID on Cardmarket’s API, also known as the 'idProduct'.
	CardmarketID *int64 `json:"cardmarket_id,omitempty"`

	// A unique ID for this card’s oracle identity. This value is consistent across reprinted card editions, and unique among different cards with the same name.
	OracleID *uuid.UUID `json:"oracle_id,omitempty"`

	// A link to where you can begin paginating all re/prints for this card on Scryfall’s API.
	PrintsSearchURI string `json:"prints_search_uri"`

	// A link to this card’s rulings list on Scryfall’s API.
	RulingsURI string `json:"rulings_uri"`

	// A link to this card’s permapage on Scryfall’s website.
	ScryfallURI string `json:"scryfall_uri"`

	// A link to this card object on Scryfall’s API.
	URI string `json:"uri"`

	// If this card is closely related to other cards, this property will be an array with Related Card Objects.
	AllParts []RelatedCard `json:"all_parts,omitempty"`

	// An array of Card Face objects, if this card is multifaced.
	CardFaces []CardFace `json:"card_faces,omitempty"`

	// The card’s converted mana cost. Note that some funny cards have fractional mana costs.
	Cmc *float64 `json:"cmc,omitempty"`

	// This card’s color identity.
	ColorIdentity []string `json:"color_identity"`

	// The colors in this card’s color indicator, if any. A null value for this field indicates the card does not have one.
	ColorIndicator []string `json:"color_indicator,omitempty"`

	// This card’s colors, if the overall card has colors defined by the rules. Otherwise the colors will be on the card_faces objects.
	Colors []string `json:"colors,omitempty"`

	// This card’s overall rank/popularity on EDHREC. Not all cards are ranked.
	EdhrecRank *int64 `json:"edhrec_rank,omitempty"`

	// This card’s hand modifier, if it is Vanguard card. This value will contain a delta, such as -1.
	HandModifier string `json:"hand_modifier,omitempty"`

	// An array of keywords that this card uses, such as 'Flying' and 'Cumulative upkeep'.
	Keywords []string `json:"keywords"`

	// A code for this card’s layout.
	Layout string `json:"layout"`

	// An object describing the legality of this card across play formats.
	Legalities map[string]string `json:"legalities"`

	// This card’s life modifier, if it is Vanguard card. This value will contain a delta, such as +2.
	LifeModifier string `json:"life_modifier,omitempty"`

	// This loyalty if any. Note that some cards have loyalties that are not numeric, such as X.
	Loyalty string `json:"loyalty,omitempty"`

	// The mana cost for this card. This value will be any empty string "" if the cost is absent.
	// Remember that per the game rules, a missing mana cost and a mana cost of {0} are different values.
	ManaCost string `json:"mana_cost,omitempty"`

	// The name of this card. If this card has multiple faces, this field will contain both names separated by ␣//␣.
	Name string `json:"name"`

	// The Oracle text for this card, if any.
	OracleText string `json:"oracle_text,omitempty"`

	// True if this card is oversized.
	Oversized bool `json:"oversized"`

	// This card’s rank/popularity on Penny Dreadful. Not all cards are ranked.
	PennyRank int64 `json:"penny_rank,omitempty"`

	// This card’s power, if any. Note that some cards have powers that are not numeric, such as *.
	Power string `json:"power,omitempty"`

	// Colors of mana that this card could produce.
	ProducedMana []string `json:"produced_mana,omitempty"`

	// True if this card is on the Reserved List.
	Reserved bool `json:"reserved"`

	// This card’s toughness, if any. Note that some cards have toughnesses that are not numeric, such as *.
	Toughness string `json:"toughness,omitempty"`

	// The type line of this card.
	TypeLine string `json:"type_line,omitempty"`

	// The name of the illustrator of this card. Newly spoiled cards may not have this field yet.
	Artist string `json:"artist,omitempty"`

	// The lit Unfinity attractions lights on this card, if any.
	AttractionLights []int64 `json:"attraction_lights,omitempty"`

	// Whether this card is found in boosters.
	Booster bool `json:"booster"`

	// This card’s border color.
	BorderColor string `json:"border_color"`

	// The Scryfall ID for the card back design present on this card.
	CardBackID *uuid.UUID `json:"card_back_id"`

	// This card’s collector number.
	CollectorNumber string `json:"collector_number"`

	// True if you should consider avoiding use of this print downstream.
	ContentWarning bool `json:"content_warning,omitempty"`

	// True if this card was only released in a video game.
	Digital bool `json:"digital"`

	// An array of computer-readable flags that indicate if this card can come in foil, nonfoil, etched, or glossy finishes.
	Finishes []string `json:"finishes"`

	// The just-for-fun name printed on the card (such as for Godzilla series cards).
	FlavorName string `json:"flavor_name,omitempty"`

	// The flavor text, if any.
	FlavorText string `json:"flavor_text,omitempty"`

	// This card’s frame effects, if any.
	FrameEffects []string `json:"frame_effects,omitempty"`

	// This card’s frame layout.
	Frame string `json:"frame"`

	// True if this card’s artwork is larger than normal.
	FullArt bool `json:"full_art"`

	// A list of games that this card print is available in, paper, arena, and/or mtgo.
	Games []string `json:"games"`

	// True if this card’s imagery is high resolution.
	HighresImage bool `json:"highres_image"`

	// A unique identifier for the card artwork that remains consistent across reprints. Newly spoiled cards may not have this field yet.
	IllustrationID *uuid.UUID `json:"illustration_id,omitempty"`

	// A computer-readable indicator for the state of this card’s image.
	ImageStatus string `json:"image_status"`

	// An object listing available imagery for this card.
	ImageURIs map[string]string `json:"image_uris,omitempty"`

	// An object containing daily price information for this card, including usd, usd_foil, usd_etched, eur, and tix prices.
	Prices *Prices `json:"prices"`

	// The localized name printed on this card, if any.
	PrintedName string `json:"printed_name,omitempty"`

	// The localized text printed on this card, if any.
	PrintedText string `json:"printed_text,omitempty"`

	// The localized type line printed on this card, if any.
	PrintedTypeLine string `json:"printed_type_line,omitempty"`

	// True if this card is a promotional print.
	Promo bool `json:"promo"`

	// An array of strings describing what categories of promo cards this card falls into.
	PromoTypes []string `json:"promo_types,omitempty"`

	// An object providing URIs to this card’s listing on major marketplaces.
	PurchaseURIs map[string]string `json:"purchase_uris,omitempty"`

	// This card’s rarity. One of common, uncommon, rare, special, mythic, or bonus.
	Rarity string `json:"rarity"`

	// An object providing URIs to this card’s listing on other Magic: The Gathering online resources.
	RelatedURIs map[string]string `json:"related_uris"`

	// The date this card was first released.
	ReleasedAt civil.Date `json:"released_at"`

	// True if this card is a reprint.
	Reprint bool `json:"reprint"`

	// A link to this card’s set on Scryfall’s website.
	ScryfallSetURI string `json:"scryfall_set_uri"`

	// This card’s full set name.
	SetName string `json:"set_name"`

	// A link to where you can begin paginating this card’s set on the Scryfall API.
	SetSearchURI string `json:"set_search_uri"`

	// The type of set this printing is in.
	SetType string `json:"set_type"`

	// A link to this card’s set object on Scryfall’s API.
	SetURI string `json:"set_uri"`

	// This card’s set code.
	Set string `json:"set"`

	// This card’s Set object UUID.
	SetID uuid.UUID `json:"set_id"`

	// True if this card is a Story Spotlight.
	StorySpotlight bool `json:"story_spotlight"`

	// True if the card is printed without text.
	Textless bool `json:"textless"`

	// Whether this card is a variation of another printing.
	Variation bool `json:"variation"`

	// The printing ID of the printing this card is a variation of.
	VariationOf string `json:"variation_of,omitempty"`

	// The security stamp on this card, if any.
	SecurityStamp string `json:"security_stamp,omitempty"`

	// This card’s watermark, if any.
	Watermark string `json:"watermark,omitempty"`

	// When, where and by whom this card was previewed.
	Preview *Preview `json:"preview,omitempty"`

	// Undocumented by Scryfall fields
	ArtistIDs []uuid.UUID `json:"artist_ids,omitempty"`
	Foil      bool        `json:"foil"`
	Nonfoil   bool        `json:"nonfoil"`
}

func (c *ScryfallClient) GetCardById(ctx context.Context, id uuid.UUID) (*Card, error) {
	urlString, err := url.JoinPath("cards", id.String())
	if err != nil {
		return nil, err
	}
	var card Card

	slog.Debug("requesting card by id", "id", id)
	resp, err := c.r(ctx).SetResult(&card).Get(urlString)
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, fmt.Errorf("failed to fetch card by id %s: %v", id, resp.Error().(Error).Details)
	}

	return &card, nil
}
