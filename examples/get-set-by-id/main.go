package main

import (
	"log"

	tcg "github.com/PokemonTCG/pokemon-tcg-sdk-go-v2/pkg"
)

func main() {
	c := tcg.NewClient("your_api_key")

	set, err := c.GetSetByID("sm115")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%s: %s\n", set.ID, set.Name)
}
