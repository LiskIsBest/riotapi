package riotapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type AccountDto struct {
	Puuid    string `json:"puuid"`
	GameName string `json:"gameName"`
	TagLine  string `json:"tagLine"`
}

type ActiveShardDto struct {
	Puuid       string `json:"puuid"`
	Game        string `json:"game"`
	ActiveShard string `json:"activeShard"`
}

type AccountsService struct {
	client *RiotApi
}

func (s AccountsService) AccountByPuuid(puuid string) (AccountDto, error) {
	endpoint := fmt.Sprintf("%s/riot/account/v1/accounts/by-puuid/%s", baseHostURL2, puuid)
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return AccountDto{}, err
	}
	req.Header.Add("X-Riot-Token", s.client.key)
	body, err := s.client.Do(req)
	if err != nil {
		return AccountDto{}, err
	}
	var account AccountDto
	err = json.Unmarshal(body, &account)
	if err != nil {
		return AccountDto{}, err
	}
	return account, nil
}

func (s AccountsService) AccountByRiotID(gameName string, tagLine string) (AccountDto, error) {
	endpoint := fmt.Sprintf("%s/riot/account/v1/accounts/by-riot-id/%s/%s", baseHostURL2, gameName, tagLine)
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return AccountDto{}, err
	}
	req.Header.Add("X-Riot-Token", s.client.key)
	body, err := s.client.Do(req)
	if err != nil {
		return AccountDto{}, err
	}
	var account AccountDto
	err = json.Unmarshal(body, &account)
	if err != nil {
		return AccountDto{}, err
	}
	return account, nil
}
