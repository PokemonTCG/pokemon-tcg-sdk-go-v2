package main

import (
	"log"

	tcg "github.com/PokemonTCG/pokemon-tcg-sdk-go-v2/pkg"
	"github.com/PokemonTCG/pokemon-tcg-sdk-go-v2/pkg/request"
)

func main() {
	c := tcg.NewClient("your_api_key")

	// Refer to https://docs.pokemontcg.io/#api_v2sets_list for how queries work
	sets, err := c.GetSets(
		request.Query("legalities.standard:legal"),
		request.PageSize(5),
		request.Page(2),
	)
	if err != nil {
		log.Fatal(err)
	}

	for _, set := range sets {
		log.Printf("%s: %s\n", set.ID, set.Name)
	}
}
