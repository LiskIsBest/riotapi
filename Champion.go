package riotapi

import (
	"encoding/json"
	"fmt"
)

type ChampionInfo struct {
	MaxNewPlayerLevel            int   `json:"maxNewPlayerLevel"`
	FreeChampionIdsForNewPlayers []int `json:"freeChampionIdsForNewPlayers"`
	FreeChampionIds              []int `json:"freeChampionIds"`
}

type ChampionService struct {
	client *RiotApi
}

// Returns champion rotations, including free-to-play and low-level free-to-play rotations
//
// https://developer.riotgames.com/apis#champion-v3/GET_getChampionInfo
func (s ChampionService) ChampionRotations() (ChampionInfo, error) {
	endpoint := fmt.Sprintf("%s/lol/platform/v3/champion-rotations", s.client.PlatformUrl)
	req, err := s.client.newRequest("GET", endpoint, nil)
	if err != nil {
		return ChampionInfo{}, err
	}
	body, err := s.client.do(req)
	if err != nil {
		return ChampionInfo{}, err
	}
	var champion ChampionInfo
	err = json.Unmarshal(body, &champion)
	if err != nil {
		return ChampionInfo{}, err
	}
	return champion, nil
}
