package main

import (
	"encoding/json"
)

type TransactionFilter struct {
	Account_Token string `json:"account_token"`
	Card_Token    string `json:"card_token"`
	Result       Result `json:"result"`
}

func TransactionListGet(config Config, filter TransactionFilter, pagination Pagination) (TransactionList, error) {
	url := config.URLBase + "transactions"
	url += Parameterize(filter, true)
	url += ParameterizePagination(pagination)

	body, err := Get(config, url)
	if err != nil { return TransactionList{}, err }

	var response TransactionList
	json.Unmarshal(body, &response)

	return response, nil
}
