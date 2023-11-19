package pokemontcgv2

import (
	"encoding/json"
	"fmt"

	"github.com/PokemonTCG/pokemon-tcg-sdk-go-v2/pkg/request"
)

// A PokemonCard represents a Pokemon card and its data.
type PokemonCard struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	Supertype    string   `json:"supertype"`
	Subtypes     []string `json:"subtypes"`
	Level        string   `json:"level"`
	Hp           string   `json:"hp"`
	Types        []string `json:"types"`
	EvolvesFrom  string   `json:"evolvesFrom"`
	EvolvesTo    []string `json:"evolvesTo"`
	Rules        []string `json:"rules"`
	AncientTrait *struct {
		Name string `json:"name"`
		Text string `json:"text"`
	} `json:"ancientTrait"`
	Abilities []struct {
		Name string `json:"name"`
		Text string `json:"text"`
		Type string `json:"type"`
	} `json:"abilities"`
	Attacks []struct {
		Name                string   `json:"name"`
		Cost                []string `json:"cost"`
		ConvertedEnergyCost int      `json:"convertedEnergyCost"`
		Damage              string   `json:"damage"`
		Text                string   `json:"text"`
	} `json:"attacks"`
	Weaknesses []struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"weaknesses"`
	Resistances []struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"resistances"`
	RetreatCost          []string `json:"retreatCost"`
	ConvertedRetreatCost int      `json:"convertedRetreatCost"`
	Set                  struct {
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
	} `json:"set"`
	Number                 string `json:"number"`
	Artist                 string `json:"artist"`
	Rarity                 string `json:"rarity"`
	FlavorText             string `json:"flavorText"`
	NationalPokedexNumbers []int  `json:"nationalPokedexNumbers"`
	Legalities             struct {
		Unlimited string `json:"unlimited"`
		Standard  string `json:"standard"`
		Expanded  string `json:"expanded"`
	} `json:"legalities"`
	Images struct {
		Small string `json:"small"`
		Large string `json:"large"`
	} `json:"images"`
	TCGPlayer struct {
		URL       string `json:"url"`
		UpdatedAt string `json:"updatedAt"`
		Prices    struct {
			Holofoil *struct {
				Low    float64 `json:"low"`
				Mid    float64 `json:"mid"`
				High   float64 `json:"high"`
				Market float64 `json:"market"`
			} `json:"holofoil,omitempty"`
			ReverseHolofoil *struct {
				Low    float64 `json:"low"`
				Mid    float64 `json:"mid"`
				High   float64 `json:"high"`
				Market float64 `json:"market"`
			} `json:"reverseHolofoil,omitempty"`
			Normal *struct {
				Low    float64 `json:"low"`
				Mid    float64 `json:"mid"`
				High   float64 `json:"high"`
				Market float64 `json:"market"`
			} `json:"normal,omitempty"`
		} `json:"prices"`
	} `json:"tcgplayer"`
	CardMarket struct {
		URL       string `json:"url"`
		UpdatedAt string `json:"updatedAt"`
		Prices    struct {
			AverageSellPrice *float64 `json:"averageSellPrice"`
			LowPrice         *float64 `json:"lowPrice"`
			TrendPrice       *float64 `json:"trendPrice"`
			GermanProLow     *float64 `json:"germanProLow"`
			SuggestedPrice   *float64 `json:"suggestedPrice"`
			ReverseHoloSell  *float64 `json:"reverseHoloSell"`
			ReverseHoloLow   *float64 `json:"reverseHoloLow"`
			ReverseHoloTrend *float64 `json:"reverseHoloTrend"`
			LowPriceExPlus   *float64 `json:"lowPriceExPlus"`
			Avg1             *float64 `json:"avg1"`
			Avg7             *float64 `json:"avg7"`
			Avg30            *float64 `json:"avg30"`
			ReverseHoloAvg1  *float64 `json:"reverseHoloAvg1"`
			ReverseHoloAvg7  *float64 `json:"reverseHoloAvg7"`
			ReverseHoloAvg30 *float64 `json:"reverseHoloAvg30"`
		} `json:"prices"`
	} `json:"cardmarket"`
}

// GetCards allows you to search and filter for cards using given options.
// Docs: https://docs.pokemontcg.io/#api_v2cards_list
func (c *apiClient) GetCards(o ...request.Option) ([]*PokemonCard, error) {
	r := request.New(endpointCards, o...)

	u, err := r.GetURL()
	if err != nil {
		return nil, err
	}

	resp, err := c.get(u)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	cardList := struct {
		Data []*PokemonCard `json:"data"`
	}{}
	if err := json.NewDecoder(resp.Body).Decode(&cardList); err != nil {
		return nil, err
	}

	return cardList.Data, nil
}

// GetCardByID returns a single pokemon card.
// Docs: https://docs.pokemontcg.io/#api_v2cards_get
func (c *apiClient) GetCardByID(id string) (*PokemonCard, error) {
	r := request.New(fmt.Sprintf(endpointCard, id))
	u, err := r.GetURL()
	if err != nil {
		return nil, err
	}

	cards := make(map[string]PokemonCard)
	resp, err := c.get(u)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&cards); err != nil {
		return nil, err
	}

	card, ok := cards["data"]
	if !ok {
		return nil, err
	}

	return &card, nil
}
