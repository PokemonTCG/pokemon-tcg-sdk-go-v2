package pokemontcgv2

import (
	"encoding/json"
	"fmt"

	"github.com/PokemonTCG/pokemon-tcg-sdk-go-v2/pkg/request"
)

// A PokemonCard represents a Pokemon card and its data.
type PokemonCard struct {
	ID           string   `json:"id" bson:"id"`
	Name         string   `json:"name" bson:"name"`
	Supertype    string   `json:"supertype" bson:"supertype"`
	Subtypes     []string `json:"subtypes" bson:"subtypes"`
	Level        string   `json:"level" bson:"level"`
	Hp           string   `json:"hp" bson:"hp"`
	Types        []string `json:"types" bson:"types"`
	EvolvesFrom  string   `json:"evolvesFrom" bson:"evolvesFrom"`
	EvolvesTo    []string `json:"evolvesTo" bson:"evolvesTo"`
	Rules        []string `json:"rules" bson:"rules"`
	AncientTrait *struct {
		Name string `json:"name" bson:"name"`
		Text string `json:"text" bson:"text"`
	} `json:"ancientTrait" bson:"ancientTrait"`
	Abilities []struct {
		Name string `json:"name" bson:"name"`
		Text string `json:"text" bson:"text"`
		Type string `json:"type" bson:"type"`
	} `json:"abilities" bson:"abilities"`
	Attacks []struct {
		Name                string   `json:"name" bson:"name"`
		Cost                []string `json:"cost" bson:"cost"`
		ConvertedEnergyCost int      `json:"convertedEnergyCost" bson:"convertedEnergyCost"`
		Damage              string   `json:"damage" bson:"damage"`
		Text                string   `json:"text" bson:"text"`
	} `json:"attacks" bson:"attacks"`
	Weaknesses []struct {
		Type  string `json:"type" bson:"type"`
		Value string `json:"value" bson:"value"`
	} `json:"weaknesses" bson:"weaknesses"`
	Resistances []struct {
		Type  string `json:"type" bson:"type"`
		Value string `json:"value" bson:"value"`
	} `json:"resistances" bson:"resistances"`
	RetreatCost          []string `json:"retreatCost" bson:"retreatCost"`
	ConvertedRetreatCost int      `json:"convertedRetreatCost" bson:"convertedRetreatCost"`
	Set                  struct {
		ID           string `json:"id" bson:"id"`
		Name         string `json:"name" bson:"name"`
		Series       string `json:"series" bson:"series"`
		PrintedTotal int    `json:"printedTotal" bson:"printedTotal"`
		Total        int    `json:"total" bson:"total"`
		Legalities   struct {
			Unlimited string `json:"unlimited" bson:"unlimited"`
			Standard  string `json:"standard" bson:"standard"`
			Expanded  string `json:"expanded" bson:"expanded"`
		} `json:"legalities" bson:"legalities"`
		PtcgoCode   string `json:"ptcgoCode" bson:"ptcgoCode"`
		ReleaseDate string `json:"releaseDate" bson:"releaseDate"`
		UpdatedAt   string `json:"updatedAt" bson:"updatedAt"`
		Images      struct {
			Symbol string `json:"symbol" bson:"symbol"`
			Logo   string `json:"logo" bson:"logo"`
		} `json:"images" bson:"images"`
	} `json:"set" bson:"set"`
	Number                 string `json:"number" bson:"number"`
	Artist                 string `json:"artist" bson:"artist"`
	Rarity                 string `json:"rarity" bson:"rarity"`
	FlavorText             string `json:"flavorText" bson:"flavorText"`
	NationalPokedexNumbers []int  `json:"nationalPokedexNumbers" bson:"nationalPokedexNumbers"`
	Legalities             struct {
		Unlimited string `json:"unlimited" bson:"unlimited"`
		Standard  string `json:"standard" bson:"standard"`
		Expanded  string `json:"expanded" bson:"expanded"`
	} `json:"legalities" bson:"legalities"`
	Images struct {
		Small string `json:"small" bson:"small"`
		Large string `json:"large" bson:"large"`
	} `json:"images" bson:"images"`
	TCGPlayer *struct {
		URL       string `json:"url" bson:"url"`
		UpdatedAt string `json:"updatedAt" bson:"updatedAt"`
		Prices    struct {
			Holofoil *struct {
				Low    float64 `json:"low" bson:"low"`
				Mid    float64 `json:"mid" bson:"mid"`
				High   float64 `json:"high" bson:"high"`
				Market float64 `json:"market" bson:"market"`
			} `json:"holofoil,omitempty" bson:"holofoil,omitempty"`
			ReverseHolofoil *struct {
				Low    float64 `json:"low" bson:"low"`
				Mid    float64 `json:"mid" bson:"mid"`
				High   float64 `json:"high" bson:"high"`
				Market float64 `json:"market" bson:"market"`
			} `json:"reverseHolofoil,omitempty" bson:"reverseHolofoil,omitempty"`
			Normal *struct {
				Low    float64 `json:"low" bson:"low"`
				Mid    float64 `json:"mid" bson:"mid"`
				High   float64 `json:"high" bson:"high"`
				Market float64 `json:"market" bson:"market"`
			} `json:"normal,omitempty" bson:"normal,omitempty"`
		} `json:"prices" bson:"prices"`
	} `json:"tcgplayer,omitempty" bson:"tcgplayer,omitempty"`
	CardMarket *struct {
		URL       string `json:"url" bson:"url"`
		UpdatedAt string `json:"updatedAt" bson:"updatedAt"`
		Prices    struct {
			AverageSellPrice *float64 `json:"averageSellPrice" bson:"averageSellPrice"`
			LowPrice         *float64 `json:"lowPrice" bson:"lowPrice"`
			TrendPrice       *float64 `json:"trendPrice" bson:"trendPrice"`
			GermanProLow     *float64 `json:"germanProLow" bson:"germanProLow"`
			SuggestedPrice   *float64 `json:"suggestedPrice" bson:"suggestedPrice"`
			ReverseHoloSell  *float64 `json:"reverseHoloSell" bson:"reverseHoloSell"`
			ReverseHoloLow   *float64 `json:"reverseHoloLow" bson:"reverseHoloLow"`
			ReverseHoloTrend *float64 `json:"reverseHoloTrend" bson:"reverseHoloTrend"`
			LowPriceExPlus   *float64 `json:"lowPriceExPlus" bson:"lowPriceExPlus"`
			Avg1             *float64 `json:"avg1" bson:"avg1"`
			Avg7             *float64 `json:"avg7" bson:"avg7"`
			Avg30            *float64 `json:"avg30" bson:"avg30"`
			ReverseHoloAvg1  *float64 `json:"reverseHoloAvg1" bson:"reverseHoloAvg1"`
			ReverseHoloAvg7  *float64 `json:"reverseHoloAvg7" bson:"reverseHoloAvg7"`
			ReverseHoloAvg30 *float64 `json:"reverseHoloAvg30" bson:"reverseHoloAvg30"`
		} `json:"prices" bson:"prices"`
	} `json:"cardmarket,omitempty" bson:"cardmarket,omitempty"`
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
