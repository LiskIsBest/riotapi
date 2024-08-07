package riotapi

import (
	"encoding/json"
	"fmt"
)

type SummonerDto struct {
	AccountId string `json:"accountId"`
	ProfileIconId int `json:"profileIconId"`
	RevisionDate int `json:"revisionDate"`
	Id int `json:"id"`
	Puuid string `json:"puuid"`
	SummonerLevel int `json:"summonerLevel"`
}

type SummonerService struct{
	client *RiotApi
}

func (s SummonerService) SummonerByPuuid(puuid string) (SummonerDto, error){
	endpoint := fmt.Sprintf("%s/lol/summoner/v4/summoners/by-puuid/%s", s.client.PlatformUrl, puuid)
	req, err := s.client.NewRequest("GET", endpoint, nil)
	if err != nil {
		return SummonerDto{}, err
	}
	body, err := s.client.Do(req)
	if err != nil {
		return SummonerDto{}, err
	}
	var summoner SummonerDto
	err = json.Unmarshal(body, &summoner)
	if err != nil{
		return SummonerDto{}, err
	}
	return summoner, nil
}