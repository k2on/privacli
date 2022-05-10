package main

import (
	"encoding/json"
)


func CardListGet(config Config) ([]Card, error) {
	url := config.URLBase + "cards?page=1&page_size=50"
	body, err := Get(config, url)
	if err != nil { return []Card{}, err }

	var response CardList
    json.Unmarshal(body, &response)
	return response.Cards, err
}
