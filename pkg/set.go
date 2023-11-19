package pokemontcgv2

import (
	"encoding/json"
	"fmt"

	"github.com/PokemonTCG/pokemon-tcg-sdk-go-v2/pkg/request"
)

// A Set is a set of cards, e.g. Base Set or Sword & Shield: Vivid Voltage.
type Set struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Series       string `json:"series"`
	PrintedTotal int    `json:"printedTotal"`
	Total        int    `json:"total"`
	Legalities   struct {
		Unlimited string `json:"unlimited"`
		Standard  string `json:"standard"`
		Expanded  string `json:"expanded"`
	} `json:"legalities"`
	PtcgoCode   string `json:"ptcgoCode"`
	ReleaseDate string `json:"releaseDate"`
	UpdatedAt   string `json:"updatedAt"`
	Images      struct {
		Symbol string `json:"symbol"`
		Logo   string `json:"logo"`
	} `json:"images"`
}

// GetSets allows you to search and filter for sets using given options.
// Docs: https://docs.pokemontcg.io/#api_v2sets_list
func (c *apiClient) GetSets(o ...request.Option) ([]*Set, error) {
	r := request.New(endpointSets, o...)
	u, err := r.GetURL()
	if err != nil {
		return nil, err
	}

	resp, err := c.get(u)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	setList := struct {
		Data []*Set `json:"data"`
	}{}
	if err := json.NewDecoder(resp.Body).Decode(&setList); err != nil {
		return nil, err
	}

	return setList.Data, nil
}

// GetSetByID returns a single set.
// Docs: https://docs.pokemontcg.io/#api_v2sets_get
func (c *apiClient) GetSetByID(id string) (*Set, error) {
	r := request.New(fmt.Sprintf(endpointSet, id))
	u, err := r.GetURL()
	if err != nil {
		return nil, err
	}

	sets := make(map[string]Set)
	resp, err := c.get(u)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&sets); err != nil {
		return nil, err
	}

	set, ok := sets["data"]
	if !ok {
		return nil, err
	}

	return &set, nil
}
