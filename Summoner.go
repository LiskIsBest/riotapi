package riotapi

import (
	"encoding/json"
	"errors"
	"fmt"
)

type SummonerDto struct {
	ProfileIconId int    `json:"profileIconId"`
	RevisionDate  int    `json:"revisionDate"`
	Puuid         string `json:"puuid"`
	SummonerLevel int    `json:"summonerLevel"`
}

type SummonerService struct {
	client *RiotApi
}

// Get a summoner by its RSO encrypted PUUID
//
// https://developer.riotgames.com/apis#summoner-v4/GET_getByRSOPUUID
func (s SummonerService) SummonerByPuuid(puuid string) (SummonerDto, error) {
	endpoint := fmt.Sprintf("%s/lol/summoner/v4/summoners/by-puuid/%s", s.client.PlatformUrl, puuid)
	req, err := s.client.newRequest("GET", endpoint, nil)
	if err != nil {
		return SummonerDto{}, err
	}
	body, err := s.client.do(req)
	if err != nil {
		return SummonerDto{}, err
	}
	var summoner SummonerDto
	err = json.Unmarshal(body, &summoner)
	if err != nil {
		return SummonerDto{}, err
	}
	return summoner, nil
}

// Get a summoner by access token
//
// https://developer.riotgames.com/apis#summoner-v4/GET_getByAccessToken
func (s SummonerService) SummonerByAccessToken(accessToken string) (SummonerDto, error){
	return SummonerDto{}, errors.New("Summoner.Summoners not implemented yet")
}