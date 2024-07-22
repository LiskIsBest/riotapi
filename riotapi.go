package riotapi

import (
	"io"
	"net/http"
)

/*
link for getting a key: https://developer.riotgames.com/
enpoints: https://developer.riotgames.com/apis
riot docs:
	- https://developer.riotgames.com/docs/portal#_getting-started
	- https://developer.riotgames.com/docs/lol
*/

var baseHostURL1 string = "https://na1.api.riotgames.com"
var baseHostURL2 string = "https://americas.api.riotgames.com"

type RiotApi struct {
	client *http.Client

	key      string
	Accounts *AccountsService
}

func NewConnection(apiKey string) *RiotApi {
	api := &RiotApi{}
	client := http.Client{}
	
	api.client = &client
	api.key = apiKey
	api.Accounts = &AccountsService{client: api}
	return api
}

func (r RiotApi) Do(req *http.Request) ([]byte, error) {
	resp, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
