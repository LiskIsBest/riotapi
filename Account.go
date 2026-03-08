package riotapi

import (
	"encoding/json"
	"fmt"
	"errors"
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

type AccountRegionDTO struct {
	Puuid  string `json:"puuid"`
	Game   string `json:"game"`
	Region string `json:"region"`
}

type AccountService struct {
	client *RiotApi
}

// Get account username, tag, and puuid by puuid
//
// https://developer.riotgames.com/apis#account-v1/GET_getByPuuid
func (s AccountService) GetByPuuid(puuid string) (AccountDto, error) {
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
func (s AccountService) GetByRiotId(gameName string, tagLine string) (AccountDto, error) {
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
func (s AccountService) GetActiveShard(game string, puuid string) (ActiveShardDto, error) {
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

// Get active region (lol and tft)
// 
// Game options: "lol" and "tft"
//
// https://developer.riotgames.com/apis#account-v1/GET_getActiveRegion
func (s AccountService) GetActiveRegion(game string, puuid string) (AccountRegionDTO, error) {
	return AccountRegionDTO{}, errors.New("Account.GetActiveRegion not implemented yet.")
}

// Get account by access token
//
// https://developer.riotgames.com/apis#account-v1/GET_getByAccessToken
func (s AccountService) GetByAccessToken(token string) (AccountDto, error) {
	return AccountDto{}, errors.New("Account.GetByAccessToken not implemented yet.")
}