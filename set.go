package scryfall

import (
	"context"
	"fmt"

	"cloud.google.com/go/civil"
)

// Set represents a group of related Magic cards. All Card objects on Scryfall
// belong to exactly one set.
//
// Due to Magic’s long and complicated history, Scryfall includes many
// un-official sets as a way to group promotional or outlier cards together.
// Such sets will likely have a code that begins with `p` or `t`, such as `pcel`
// or `tori`.
//
// Official sets always have a three-letter set code, such as `zen`.
type Set struct {
	// A content type for this object, always "set".
	ObjectType string `json:"object"`

	// A unique ID for this set on Scryfall that will not change.
	ID string `json:"id"`

	// The unique three to five-letter code for this set.
	Code string `json:"code"`

	// The unique code for this set on MTGO, which may differ from the regular code.
	MtgoCode string `json:"mtgo_code,omitempty"`

	// This set’s ID on TCGplayer’s API, also known as the groupId.
	TcgplayerID *int `json:"tcgplayer_id,omitempty"`

	// The English name of the set.
	Name string `json:"name"`

	// A computer-readable classification for this set.
	SetType string `json:"set_type"`

	// The date the set was released or the first card was printed in the set.
	ReleasedAt civil.Date `json:"released_at,omitempty"`

	// The block code for this set, if any.
	BlockCode string `json:"block_code,omitempty"`

	// The block or group name code for this set, if any.
	Block string `json:"block,omitempty"`

	// The set code for the parent set, if any. promo and token sets often have a parent set.
	ParentSetCode string `json:"parent_set_code,omitempty"`

	// The number of cards in this set.
	CardCount int `json:"card_count"`

	// The denominator for the set’s printed collector numbers.
	PrintedSize *int `json:"printed_size,omitempty"`

	// True if this set was only released in a video game.
	Digital bool `json:"digital"`

	// True if this set contains only foil cards.
	FoilOnly bool `json:"foil_only"`

	// True if this set contains only nonfoil cards.
	NonfoilOnly bool `json:"nonfoil_only"`

	// A link to this set’s permapage on Scryfall’s website.
	ScryfallURI string `json:"scryfall_uri"`

	// A link to this set object on Scryfall’s API.
	URI string `json:"uri"`

	// A URI to an SVG file for this set’s icon on Scryfall’s CDN.
	IconSvgURI string `json:"icon_svg_uri"`

	// A Scryfall API URI that you can request to begin paginating over the cards in this set.
	SearchURI string `json:"search_uri"`

	// The unique code for this set on MTG Arena, which may differ from the regular code.
	ArenaCode string `json:"arena_code,omitempty"`
}

func (c *ScryfallClient) PaginateAllSets(ctx context.Context, callback PaginationCallback[Set]) error {
	var list List[Set]
	resp, err := c.r(ctx).
		SetResult(&list).
		Get("sets")
	if err != nil {
		return fmt.Errorf("failed to fetch list of sets: %v", err)
	}
	if !resp.IsSuccess() {
		return fmt.Errorf("failed to fetch list of sets: %s", resp.Status())
	}
	return list.Paginate(ctx, c, callback)
}
