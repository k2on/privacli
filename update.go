package main

import (
	"encoding/json"
)

type OptionsUpdate struct {
	Memo               string   `json:"memo"`
	SpendLimit         int      `json:"spend_limit"`
	SpendLimitDuration Duration `json:"spend_limit_duration"`
	State              State    `json:"stat"`
	Pin                string   `json:"pin"`
}

func CardUpdate(config Config, id string, options OptionsUpdate) (Card, error) {
	url := config.URLBase + "cards/" + id

	bytes, err := json.Marshal(options)
	if err != nil { return Card{}, err }
	data := string(bytes)

	body, err := Patch(config, url, data)
	if err != nil { return Card{}, err }

	var response Card
    json.Unmarshal(body, &response)

	return response, nil
}

func CardClose(config Config, id string) (Card, error) {
	return CardUpdate(config, id, OptionsUpdate{
		State: StateClosed,
	})
}