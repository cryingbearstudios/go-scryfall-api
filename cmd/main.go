package main

import (
	"context"
	"log"

	scryfall "cryingbear.net/go-scryfall-api"
	"resty.dev/v3"
)

func main() {
	ctx := context.Background()
	client := scryfall.NewClient(ctx)
	defer client.Close()

	_, err := client.GetBulkDataByType("all_cards")
	if err != nil {
		log.Fatalf("failed to look up all_cards bulk data info: %v", err)
	}

	restyClient := resty.New()
	defer restyClient.Close()

}
