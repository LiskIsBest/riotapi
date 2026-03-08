package riotapi

import (
	"errors"
)

type LeagueExpService struct {
	client *RiotApi
}

// Get the master league for given queue
//
// https://developer.riotgames.com/apis#league-v4/GET_getMasterLeague
func (s LeagueExpService) Entries(queue string, tier string, division string)(string, error){
	return "", errors.New("LeagueExp.Entries not implemented yet.")
}

