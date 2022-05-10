package main

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

type Method string
const (
	MethodGet Method = "GET"
	MethodPost       = "POST"
	MethodPatch      = "PATCH"
)

func Request(config Config, url string, method Method, data string) ([]byte, error) {
	request, err := MakeRequest(method, url, data)
	if err != nil { return nil, err }
	isPayload := data != ""
	return ExecRequest(request, config.key, isPayload)
}

func MakeRequest(method Method, url string, data string) (*http.Request, error) {
	reader := strings.NewReader(data)
	return http.NewRequest(string(method), url, reader)
}

func ExecRequest(request *http.Request, key string, isPayload bool) ([]byte, error) {
	AddHeaders(request, key, isPayload)
	res, err := http.DefaultClient.Do(request)
	if err != nil { return nil, err }

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil { return nil, err }

	if res.Status != "200 OK" { return nil, errors.New(string(body)) }
	return body, nil
}

func AddHeaders(request *http.Request, key string, hasContent bool) {
	request.Header.Add("Accept", "application/json")
	request.Header.Add("Authorization", key)
	if !hasContent { return }
	request.Header.Add("Content-Type", "application/json")
}

func Get(config Config, url string) ([]byte, error) {
	return Request(config, url, MethodGet, "")
}

func Post(config Config, url string, payload string) ([]byte, error) {
	return Request(config, url, MethodPost, payload)
}

func Patch(config Config, url string, payload string) ([]byte, error) {
	return Request(config, url, MethodPatch, payload)
}