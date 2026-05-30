package scryfall

import "github.com/govalues/decimal"

type Prices struct {
	Usd       *decimal.Decimal `json:"usd"`
	UsdFoil   *decimal.Decimal `json:"usd_foil"`
	UsdEtched *decimal.Decimal `json:"usd_etched"`
	Eur       *decimal.Decimal `json:"eur"`
	EurFoil   *decimal.Decimal `json:"eur_foil"`
	Tix       *decimal.Decimal `json:"tix"`
}
