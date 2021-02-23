package pokemontcgv2

import "testing"

func TestGetRarities(t *testing.T) {
	// Test a subset of known rarities
	wantRarities := []string{"Amazing Rare", "Common", "Rare Holo VMAX"}
	types, err := NewClient("").GetRarities()
	if err != nil {
		t.Fatal(err)
	}

	for _, want := range wantRarities {
		if !in(want, types) {
			t.Errorf("known rarity not found in API rarity: %s", want)
		}
	}
}
