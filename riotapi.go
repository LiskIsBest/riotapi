package riotapi

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

/*
link for getting a key: https://developer.riotgames.com/
enpoints: https://developer.riotgames.com/apis
riot docs:
	- https://developer.riotgames.com/docs/portal#_getting-started
	- https://developer.riotgames.com/docs/lol
*/

func convertPlatform2Regional(region string) (string, error) {
	switch strings.ToUpper(region) {
	case "NA1", "BR1", "LA1", "LA2":
		return "americas", nil
	case "KR", "JP1":
		return "asia", nil
	case "EUN1", "EUW1", "TR1", "RU":
		return "europe", nil
	case "OC1", "PH2", "SG2", "TH2", "TW2", "VN2":
		return "sea", nil
	case "AMERICAS", "ASIA", "EUROPE", "SEA":
		return region, nil
	default:
		return "", errors.New(fmt.Sprintf("Entered region: %s, is not a valid region.", region))
	}
}

var baseHostURL string = ".api.riotgames.com"

type RiotApi struct {
	client *http.Client

	key string

	Accounts *AccountsService
}

func NewConnection(region string, apiKey string) *RiotApi {
	client := http.Client{}
	return &RiotApi{client: &client, key: apiKey}
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
