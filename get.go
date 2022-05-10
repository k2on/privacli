package main

import (
	"encoding/json"
)

func CardGet(config Config, id string) (Card, error) {
	url := config.URLBase + "cards/" + id

	body, err := Get(config, url)
	if err != nil { return Card{}, err }

	var response Card
    json.Unmarshal(body, &response)

	return response, err
}
