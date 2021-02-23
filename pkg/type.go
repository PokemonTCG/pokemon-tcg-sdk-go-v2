package pokemontcgv2

import (
	"encoding/json"

	"github.com/PokemonTCG/pokemon-tcg-sdk-go-v2/pkg/request"
)

func (c *apiClient) GetTypes() ([]string, error) {
	return getTypes(c, endpointTypes)
}

func (c *apiClient) GetSuperTypes() ([]string, error) {
	return getTypes(c, endpointSuperTypes)
}

func (c *apiClient) GetSubTypes() ([]string, error) {
	return getTypes(c, endpointSubTypes)
}

// All 3 "types" return the same data type from the API.
func getTypes(c *apiClient, endpoint string) ([]string, error) {
	r := request.New(endpoint)
	u, err := r.GetURL()
	if err != nil {
		return nil, err
	}

	resp, err := c.get(u)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	types := struct {
		Data []string `json:"data"`
	}{}
	if err := json.NewDecoder(resp.Body).Decode(&types); err != nil {
		return nil, err
	}

	return types.Data, nil
}
