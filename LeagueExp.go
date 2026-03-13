package riotapi

import (
	"fmt"
	"encoding/json"
)

type LeagueExpService struct {
	client *RiotApi
}

// Get the master league for given queue
//
// https://developer.riotgames.com/apis#league-v4/GET_getMasterLeague
func (s LeagueExpService) GetMasterLeague(queue string, tier string, division string)([]LeagueEntryDto, error){
	endpoint := fmt.Sprintf("%s/lol/league-exp/v4/entries/%s/%s/%s", s.client.PlatformUrl,queue,tier,division)
	req, err := s.client.newRequest("GET", endpoint, nil)
	if err != nil {
		return []LeagueEntryDto{}, err
	}
	body, err := s.client.do(req)
	if err != nil {
		return []LeagueEntryDto{}, err
	}
	var league []LeagueEntryDto
	err = json.Unmarshal(body, &league)
	if err != nil {
		return []LeagueEntryDto{}, err
	}
	return league, nil
}

