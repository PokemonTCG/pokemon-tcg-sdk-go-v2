package pokemontcgv2

import (
	"strings"
	"testing"

	"github.com/PokemonTCG/pokemon-tcg-sdk-go-v2/pkg/request"
)

func TestGetCards(t *testing.T) {
	// Test correctness by just checking that we _exhaust_ a set of IDs for a given
	// search. Since this tests against live data, there may be a point in the future
	// where the API returns more Pokemon than expected - that's why it exhausts the set.
	// In the future where the search query actually exceeds the limit of 250 Pokemon per page,
	// well, there will be a ton of cards and this test needs to be fixed :)
	tests := []struct {
		name    string
		o       []request.Option
		wantIDs []string
	}{
		{
			name: "Psychic Jirachi",
			o: []request.Option{
				request.Query("name:jirachi", "types:psychic"),
			},
			wantIDs: []string{
				"g1-RC13",
				"sm11-79",
				"hgss2-1",
			},
		},
		{
			name: "Flareon with more than 70 HP",
			o: []request.Option{
				request.Query("name:flareon", "hp:[70 TO *]"),
			},
			wantIDs: []string{
				"pop3-2",
				"bwp-BW88",
				"swshp-SWSH041",
				"sm12-25",
			},
		},
		{
			name: "Burning shadows, trainer, ultra rare",
			o: []request.Option{
				request.Query("supertype:trainer", "set.id:sm3", "rarity:Rare Ultra"),
			},
			wantIDs: []string{
				"sm3-142",
				"sm3-143",
				"sm3-144",
				"sm3-145",
			},
		},
		{
			name: "Dash in name of Pokemon",
			o: []request.Option{
				request.Query("name:gardevoir-gx"),
			},
			wantIDs: []string{
				"sma-SV75",
				"sm3-93",
				"sm3-140",
				"sm3-159",
			},
		},
		{
			name: "Space in name of Pokemon and wildcard",
			o: []request.Option{
				request.Query("name:gardevoir &*"),
			},
			wantIDs: []string{
				"sm10-130",
				"sm10-204",
				"sm10-205",
				"sm10-225",
			},
		},
	}

	c := NewClient("")
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cards, err := c.GetCards(test.o...)
			if err != nil {
				t.Error(err)
				return
			}
			// Preprocess want IDs to a set
			wantIDs := make(map[string]interface{})
			for _, id := range test.wantIDs {
				wantIDs[id] = struct{}{}
			}
			for _, card := range cards {
				if _, ok := wantIDs[card.ID]; ok {
					delete(wantIDs, card.ID)
				}
			}

			if len(wantIDs) > 0 {
				var leftovers []string
				for k := range wantIDs {
					leftovers = append(leftovers, k)
				}
				t.Errorf("did not exhaust cards: %s", strings.Join(leftovers, ", "))
			}
		})
	}
}

func TestGetCards_LimitNumber(t *testing.T) {
	c := NewClient("")
	cards, err := c.GetCards(request.PageSize(1))
	if err != nil {
		t.Fatal(err)
	}

	if len(cards) != 1 {
		t.Errorf("want 1 card, got %d", len(cards))
	}
}

func TestGetCardByID(t *testing.T) {
	const want = "xyp-XY67a"

	c := NewClient("")
	card, err := c.GetCardByID(want)
	if err != nil {
		t.Fatal(err)
	}

	if card.ID != want {
		// Really if this test fails, something is _probably_ wrong upstream :)
		t.Errorf("wrong ID, got %s, want %s", card.ID, want)
	}
}
