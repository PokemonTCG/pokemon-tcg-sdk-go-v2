package pokemontcgv2

import "testing"

func TestTypes(t *testing.T) {
	// Test a subset of known types
	wantTypes := []string{"Fire", "Water", "Psychic", "Lightning"}
	types, err := NewClient("").GetTypes()
	if err != nil {
		t.Fatal(err)
	}

	for _, want := range wantTypes {
		if !in(want, types) {
			t.Errorf("known type not found in API types: %s", want)
		}
	}
}

func TestSuperTypes(t *testing.T) {
	// Test a (sub)set of all super types
	wantTypes := []string{"Energy", "Pokémon", "Trainer"}
	types, err := NewClient("").GetSuperTypes()
	if err != nil {
		t.Fatal(err)
	}

	for _, want := range wantTypes {
		if !in(want, types) {
			t.Errorf("known super type not found in API super types: %s", want)
		}
	}
}

func TestSubTypes(t *testing.T) {
	// Test a subset of known sub types
	wantTypes := []string{"BREAK", "GX", "LEGEND", "Stadium"}
	types, err := NewClient("").GetSubTypes()
	if err != nil {
		t.Fatal(err)
	}

	for _, want := range wantTypes {
		if !in(want, types) {
			t.Errorf("known sub type not found in API sub types: %s", want)
		}
	}
}

func in(a string, bs []string) bool {
	for _, b := range bs {
		if a == b {
			return true
		}
	}
	return false
}
