package main

import (
	"fmt"
	"io/ioutil"
)

const URL_BASE = "https://sandbox.privacy.com/v1/"

func GetToken() string {
	content, err := ioutil.ReadFile("key")
	if err != nil { panic(err) }
	return string(content)
}

func main() {
	token := GetToken()

	config := Config{
		URLBase: URL_BASE,
		key: token,
	}

	t, e := TransactionListGet(config, TransactionFilter{
		Card_Token: "3d3a293e-5cce-4203-bbfc-3992b5a43b0a",
		Result: ResultApproved,
	}, Pagination{})

	fmt.Println(t)
	fmt.Println(e)

	// cards, _ := CardList(config)
	// for _, c := range cards {
	// 	fmt.Println(c.Memo + ": " + c.Token)
	// }
}
