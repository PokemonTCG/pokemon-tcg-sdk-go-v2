package request

import (
	"strconv"
	"testing"
)

func TestQuery(t *testing.T) {
	tests := []struct {
		name  string
		query []string
		want  string
	}{
		{
			name:  "Space replacement",
			query: []string{"rarity:Rare Ultra"},
			want:  `rarity:"Rare Ultra"`,
		},
		{
			name:  "No space replacement on range",
			query: []string{"hp:[10 TO 50]"},
			want:  "hp:[10 TO 50]",
		},
		{
			name: "Mix of space replacement and none",
			query: []string{
				"rarity:Rare Ultra",
				"HP:[10 TO 50]",
			},
			want: `rarity:"Rare Ultra" HP:[10 TO 50]`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := New("", Query(test.query...))
			if r.options[queryKey] != test.want {
				t.Errorf("bad query, got %s, want %s", r.options[queryKey], test.want)
			}
		})
	}
}

func TestPage(t *testing.T) {
	tests := []struct {
		page int
		want int
	}{
		{
			page: -1,
			want: 1,
		},
		{
			page: 5,
			want: 5,
		},
		{
			page: 100,
			want: 100,
		},
	}
	for _, test := range tests {
		r := New("", Page(test.page))
		if r.options[pageKey] != strconv.Itoa(test.want) {
			t.Errorf("bad page, got %s, want %d", r.options[pageKey], test.page)
		}
	}
}

func TestPageSize(t *testing.T) {
	tests := []struct {
		pageSize int
		want     int
	}{
		{
			pageSize: 0,
			want:     1,
		},
		{
			pageSize: 5,
			want:     5,
		},
		{
			pageSize: 251,
			want:     250,
		},
	}
	for _, test := range tests {
		r := New("", PageSize(test.pageSize))
		if r.options[pageSizeKey] != strconv.Itoa(test.want) {
			t.Errorf("bad page, got %s, want %d", r.options[pageSizeKey], test.pageSize)
		}
	}
}

func TestOrderBy(t *testing.T) {
	tests := []struct {
		name string
		keys []string
		want string
	}{
		{
			name: "One key",
			keys: []string{"+name"},
			want: "+name",
		},
		{
			name: "Two keys",
			keys: []string{"+name", "-id"},
			want: "+name,-id",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := New("", OrderBy(test.keys...))
			if r.options[orderByKey] != test.want {
				t.Errorf("bad order by, got %s, want %s", r.options[orderByKey], test.want)
			}
		})
	}
}
