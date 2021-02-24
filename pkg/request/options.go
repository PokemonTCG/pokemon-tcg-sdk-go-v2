package request

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

const (
	queryKey    = "q"
	pageKey     = "page"
	pageSizeKey = "pageSize"
	orderByKey  = "orderBy"
)

// An Option can be applied to a request to modify the query.
type Option func(r *Request)

// Query sets a query for the request. Valid for GetCards and GetSets.
//
// Find examples here: https://docs.pokemontcg.io/#api_v2cards_list
// and: https://docs.pokemontcg.io/#api_v2sets_list
func Query(query ...string) func(r *Request) {
	// Take care of the case of spaces in the search query - surround it with quotes.
	// For example, `rarity:Rare Ultra` becomes `rarity:"Rare Ultra"`
	// However, if the user is using a range search e.g. hp:[10 TO 100] or already
	// includes quotes, don't surround it.
	for i, q := range query {
		spl := strings.Split(q, ":")
		if len(spl) == 2 {
			if strings.Contains(spl[1], " ") &&
				!strings.ContainsAny(spl[1], "[]\"") {
				spl[1] = fmt.Sprintf(`"%s"`, spl[1])
			}
		}
		query[i] = strings.Join(spl, ":")
	}
	return func(r *Request) {
		r.options[queryKey] = strings.Join(query, " ")
	}
}

// Page sets a page for the request. Defaults to 1.
func Page(page int) func(r *Request) {
	if page < 1 {
		log.Printf("Provided page [%d] is less than 1, defaulting to 1\n", page)
		page = 1
	}
	return func(r *Request) {
		r.options[pageKey] = strconv.Itoa(page)
	}
}

// PageSize sets a page size for the request. Defaults (and maxes) to 250.
func PageSize(pageSize int) func(r *Request) {
	if pageSize > 250 {
		log.Printf("Provided pageSize [%d] is gretaer than 250, defaulting to 250\n", pageSize)
		pageSize = 250
	}
	if pageSize < 1 {
		log.Printf("Provided pageSize [%d] is less than 1, defaulting to 1\n", pageSize)
		pageSize = 1
	}
	return func(r *Request) {
		r.options[pageSizeKey] = strconv.Itoa(pageSize)
	}
}

// OrderBy sets an order key and ascending/descending for the request.
// Find examples here: https://docs.pokemontcg.io/#api_v2cards_list
func OrderBy(keys ...string) func(r *Request) {
	return func(r *Request) {
		r.options[orderByKey] = strings.Join(keys, ",")
	}
}
