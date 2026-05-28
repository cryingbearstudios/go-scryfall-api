package scryfall_test

import (
	"context"
	"log"
	"os"
	"testing"

	"cryingbear.net/go-scryfall-api"
	"github.com/google/uuid"
	"github.com/govalues/decimal"
	"github.com/stretchr/testify/assert"
)

var AlphaCounterspell = uuid.MustParse("0df55e3f-14de-46ef-b6b1-616618724d9e")
var CounterspellOracleId = uuid.MustParse("cc187110-1148-4090-bbb8-e205694a39f5")
var loadedCard *scryfall.Card

func TestMain(m *testing.M) {
	client := scryfall.NewClient()
	var err error
	loadedCard, err = client.GetCardById(context.Background(), AlphaCounterspell)
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(m.Run())
}

func TestUuids(t *testing.T) {
	assert.Equal(t, AlphaCounterspell, loadedCard.ID)
	if assert.NotNil(t, loadedCard.OracleID) {
		assert.Equal(t, CounterspellOracleId, *loadedCard.OracleID)
	}
}

func TestDecimals(t *testing.T) {
	assert.True(t, decimal.One.Cmp(loadedCard.Prices["usd"]) < 0)
	assert.True(t, decimal.One.Cmp(loadedCard.Prices["eur"]) < 0)
}
