package pokemontcgv2

import (
	"net/http"

	"github.com/PokemonTCG/pokemon-tcg-sdk-go-v2/pkg/request"
)

const (
	apiEndpoint = "https://api.pokemontcg.io/v2/"
	// endpointCards is the endpoint to list cards.
	endpointCards = apiEndpoint + "cards"
	// endpointCard is the endpoint template to get a single card.
	endpointCard = apiEndpoint + "cards/%s"
	// endpointSets is the endpoint to list sets.
	endpointSets = apiEndpoint + "sets"
	// endpointSet is the endpoint template to get a single set.
	endpointSet = apiEndpoint + "sets/%s"
	// endpointTypes is the endpoint to get types.
	endpointTypes = apiEndpoint + "types"
	// endpointSuperTypes is the endpoint to get super types.
	endpointSuperTypes = apiEndpoint + "supertypes"
	// endpointSubTypes is the endpoint to get sub types.
	endpointSubTypes = apiEndpoint + "subtypes"
	// endpointRarities is the endpoint to get all rarities.
	endpointRarities = apiEndpoint + "rarities"
)

// A Client is the main entrypoint to the Pokemontcg.io API.
type Client interface {
	GetCards(o ...request.Option) ([]*PokemonCard, error)
	GetCardByID(id string) (*PokemonCard, error)
	GetSets(o ...request.Option) ([]*Set, error)
	GetSetByID(id string) (*Set, error)
	GetTypes() ([]string, error)
	GetSubTypes() ([]string, error)
	GetSuperTypes() ([]string, error)
	GetRarities() ([]string, error)
}

type apiClient struct {
	apiKey string
}

// NewClient creates a new client with the given API key.
//
// It is possible to create an authenticated client by passing in an empty string.
// However, these clients are restricted with higher rate limits.
func NewClient(APIKey string) Client {
	return &apiClient{APIKey}
}

func (c *apiClient) get(url string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	if len(c.apiKey) > 0 {
		req.Header.Set("X-Api-Key", c.apiKey)
	}
	return client.Do(req)
}
