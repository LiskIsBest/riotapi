package riotapi

import (
	"encoding/json"
	"fmt"
	"slices"
)

type LeagueListDto struct {
	LeagueId string          `json:"leagueId"`
	Entries  []LeagueItemDto `json:"entries"`
	Tier     string          `json:"tier"`
	Name     string          `json:"name"`
	Queue    string          `json:"queue"`
}

type LeagueItemDto struct {
	FreshBlood   bool          `json:"freshBlood"`
	Wins         int           `json:"wins"`
	MiniSeries   MiniSeriesDto `json:"miniSeries"`
	Inactive     bool          `json:"inactive"`
	Veteran      bool          `json:"veteran"`
	HotStreak    bool          `json:"hotStreak"`
	Rank         string        `json:"rank"`
	LeaguePoints int           `json:"leaguePoints"`
	Losses       int           `json:"losses"`
	SummonerId   string        `json:"summonerId"`
}

type LeagueEntryDto struct {
	LeagueId string `json:"leagueId"`
	SummonerId string `json:"summonerId"`
	QueueType string `json:"queueType"`
	Tier string `json:"tier"`
	Rank string `json:"rank"`
	LeaguePoints int `json:"leaguePoints"`
	Wins int `json:"wins"`
	Losses int `json:"losses"`
	HotStreak bool `json:"hotStreak"`
	Veteran bool `json:"veteran"`
	FreshBlood bool `json:"freshBlood"`
	Inactive bool `json:"inactive"`
	MiniSeries MiniSeriesDto `json:"miniSeries"`
}

type MiniSeriesDto struct {
	Losses   int    `json:"losses"`
	Progress string `json:"progress"`
	Target   int    `json:"target"`
	Wins     int    `json:"wins"`
}

type LeagueService struct {
	client *RiotApi
}

var QUEUE_OPTIONS = []string{"RANKED_SOLO_SR","RANKED_SOLO_TT","RANKED_FLEX_5x5"}
// Get the challenger league for given queue
//
// Valid queue options: RANKED_SOLO_SR, RANKED_SOLO_TT, RANKED_FLEX_5x5
//
// https://developer.riotgames.com/apis#league-v4/GET_getChallengerLeague
func (s LeagueService) Challengerleagues(queue string) (LeagueListDto, error) {
	if !slices.Contains(QUEUE_OPTIONS, queue) {
		return LeagueListDto{}, fmt.Errorf("queue: %s. is not a valid option", queue)
	}
	endpoint := fmt.Sprintf("%s/lol/league/v4/challengerleagues/by-queue/%s",s.client.PlatformUrl, queue)
	req, err := s.client.NewRequest("GET", endpoint, nil)
	if err != nil {
		return LeagueListDto{}, err
	}
	body, err := s.client.Do(req)
	if err != nil {
		return LeagueListDto{}, err
	}
	var leagueList LeagueListDto
	err = json.Unmarshal(body, &leagueList)
	if err != nil {
		return LeagueListDto{}, err
	}
	return leagueList, nil
}

//Get league entries in all queues for a given summoner ID
//
//https://developer.riotgames.com/apis#league-v4/GET_getLeagueEntriesForSummoner
func (s LeagueService) EntriesBySummonerId(summonerId string) (){

}