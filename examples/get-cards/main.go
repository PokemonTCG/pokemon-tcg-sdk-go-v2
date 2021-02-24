package main

import (
	"log"

	tcg "github.com/PokemonTCG/pokemon-tcg-sdk-go-v2/pkg"
	"github.com/PokemonTCG/pokemon-tcg-sdk-go-v2/pkg/request"
)

func main() {
	c := tcg.NewClient("your_api_key")

	// Refer to https://docs.pokemontcg.io/#api_v2cards_list for how queries work
	cards, err := c.GetCards(
		request.Query("name:jirachi", "types:psychic"),
		request.OrderBy("+name"),
		request.PageSize(3),
		request.Page(2),
	)
	if err != nil {
		log.Fatal(err)
	}

	for _, card := range cards {
		log.Printf("%s: %s\n", card.Name, card.Set.Name)
	}
}
