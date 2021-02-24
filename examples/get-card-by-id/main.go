package main

import (
	"log"

	tcg "github.com/PokemonTCG/pokemon-tcg-sdk-go-v2/pkg"
)

func main() {
	c := tcg.NewClient("your_api_key")

	card, err := c.GetCardByID("sm35-42")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%s: %s\n", card.Name, card.Set.Name)
}
