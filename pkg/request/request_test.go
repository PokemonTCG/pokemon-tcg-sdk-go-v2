package request

import (
	"strings"
	"testing"
)

func TestGetURL(t *testing.T) {
	const endpoint = "example.com"
	tests := []struct {
		name    string
		options []Option
		want    []string
	}{
		{
			name: "Single name query",
			options: []Option{
				Query("name:jirachi"),
			},
			want: []string{"q=name%3Ajirachi"},
		},
		{
			name: "Multiple queries",
			options: []Option{
				Query("name:jirachi", "set.id:swsh4", "types:normal"),
			},
			want: []string{"q=name%3Ajirachi+set.id%3Aswsh4+types%3Anormal"},
		},
		{
			name: "Multiple queries and page",
			options: []Option{
				Query("name:jirachi", "set.id:swsh4", "types:normal"),
				Page(100),
			},
			want: []string{
				"q=name%3Ajirachi+set.id%3Aswsh4+types%3Anormal",
				"page=1",
			},
		},
		{
			name: "Multiple queries, page, pageSize, and orderBy",
			options: []Option{
				Query("name:jirachi", "set.id:swsh4", "types:normal"),
				Page(100),
				PageSize(50),
				OrderBy("+name", "-id"),
			},
			want: []string{
				"q=name%3Ajirachi+set.id%3Aswsh4+types%3Anormal",
				"page=1",
				"pageSize=50",
				"orderBy=%2Bname%2C-id",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := New(endpoint, test.options...)
			got, err := r.GetURL()
			if err != nil {
				t.Fatal(err)
			}
			for _, want := range test.want {
				if !strings.Contains(got, want) {
					t.Errorf("GetURL wrong URL queries, got %s want %s", got, want)
				}
			}
		})
	}
}
