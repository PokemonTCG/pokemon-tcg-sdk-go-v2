package pokemontcgv2

// GetRarities gets the rarities available in the TCG, e.g. "Amazing Rare" and "Common"
func (c *apiClient) GetRarities() ([]string, error) {
	// workaround since it's the same format as types :)
	return getTypes(c, endpointRarities)
}
