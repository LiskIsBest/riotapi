package riotapi

import (
	"encoding/json"
	"fmt"
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

// Get account username, tag, and puuid by puuid
//
// https://developer.riotgames.com/apis#account-v1/GET_getByPuuid
func (s AccountsService) AccountByPuuid(puuid string) (AccountDto, error) {
	endpoint := fmt.Sprintf("%s/riot/account/v1/accounts/by-puuid/%s", s.client.RegionalUrl, puuid)
	req, err := s.client.newRequest("GET", endpoint, nil)
	if err != nil {
		return AccountDto{}, err
	}
	body, err := s.client.do(req)
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

// Get account username, tag, and puuid by username and tag
//
// https://developer.riotgames.com/apis#account-v1/GET_getByRiotId
func (s AccountsService) AccountByRiotID(gameName string, tagLine string) (AccountDto, error) {
	endpoint := fmt.Sprintf("%s/riot/account/v1/accounts/by-riot-id/%s/%s", s.client.RegionalUrl, gameName, tagLine)
	req, err := s.client.newRequest("GET", endpoint, nil)
	if err != nil {
		return AccountDto{}, err
	}
	body, err := s.client.do(req)
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

// Get active shard for a player
//
// https://developer.riotgames.com/apis#account-v1/GET_getActiveShard
func (s AccountsService) ActiveShardByGame(game string, puuid string) (ActiveShardDto, error) {
	endpoint := fmt.Sprintf("%s/riot/account/v1/active-shards/by-game/%s/by-puuid/%s", s.client.RegionalUrl, game, puuid)
	req, err := s.client.newRequest("GET", endpoint, nil)
	if err != nil {
		return ActiveShardDto{}, err
	}
	body, err := s.client.do(req)
	if err != nil {
		return ActiveShardDto{}, err
	}
	var shard ActiveShardDto
	err = json.Unmarshal(body, &shard)
	if err != nil {
		return ActiveShardDto{}, err
	}
	return shard, nil
}
