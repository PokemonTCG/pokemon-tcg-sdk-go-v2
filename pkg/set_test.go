package pokemontcgv2

import (
	"strings"
	"testing"

	"github.com/PokemonTCG/pokemon-tcg-sdk-go-v2/pkg/request"
)

func TestGetSets(t *testing.T) {
	// See TestGetCards for rationale of this style of unit/functional test.
	tests := []struct {
		name    string
		o       []request.Option
		wantIDs []string
	}{
		{
			name: "Base Set 1",
			o: []request.Option{
				request.Query("id:base1"),
			},
			wantIDs: []string{
				"base1",
			},
		},
		{
			name: "SM Series",
			o: []request.Option{
				request.Query("series:Sun & Moon"),
			},
			wantIDs: []string{
				"smp",
				"sma",
				"sm1",
				"sm2",
			},
		},
	}

	c := NewClient("")
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			sets, err := c.GetSets(test.o...)
			if err != nil {
				t.Error(err)
				return
			}
			// Preprocess want IDs to a set
			wantIDs := make(map[string]interface{})
			for _, id := range test.wantIDs {
				wantIDs[id] = struct{}{}
			}
			for _, set := range sets {
				if _, ok := wantIDs[set.ID]; ok {
					delete(wantIDs, set.ID)
				}
			}

			if len(wantIDs) > 0 {
				var leftovers []string
				for k := range wantIDs {
					leftovers = append(leftovers, k)
				}
				t.Errorf("did not exhaust sets: %s", strings.Join(leftovers, ", "))
			}
		})
	}
}

func TestGetSetByID(t *testing.T) {
	const want = "sma"

	c := NewClient("")
	set, err := c.GetSetByID(want)
	if err != nil {
		t.Fatal(err)
	}

	if set.ID != want {
		t.Errorf("wrong set, got %s, want %s", set.ID, want)
	}
}
