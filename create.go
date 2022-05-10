package main

import (
	"encoding/json"
)

type OptionsCreate struct {
	ExpMonth           string   `json:"exp_month"`
	ExpYear            string   `json:"exp_year"`
	Memo               string   `json:"memo"`
	SpendLimit         int      `json:"spend_limit"`
	SpendLimitDuration Duration `json:"spend_limit_duration"`
	State              State    `json:"state"`
	Type               Type     `json:"type"`
	Pin                string   `json:"pin"`
}

func CardCreate(config Config, options OptionsCreate) (Card, error) {
	url := config.URLBase + "cards"

	bytes, err := json.Marshal(options)
	if err != nil { return Card{}, err }
	data := string(bytes)

	body, err := Post(config, url, data)
	if err != nil { return Card{}, err }

	var response Card
    json.Unmarshal(body, &response)

	return response, nil
}