package riotapi

import (
	"encoding/json"
	"fmt"
)

type PlayerDto struct {
	SummonerId string `json:"summonerId"`
	TeamId     string `json:"teamId"`
	Position   string `json:"position"`
	Role       string `json:"role"`
}

type TeamDto struct {
	Id           string      `json:"id"`
	TournamentId int         `json:"tournamentId"`
	Name         string      `json:"name"`
	IconId       int         `json:"iconId"`
	Tier         int         `json:"tier"`
	Captain      string      `json:"captain"`
	Abbreviation string      `json:"abbreviation"`
	Players      []PlayerDto `json:"players"`
}

type TournamentDto struct {
	Id               int                  `json:"id"`
	ThemeId          int                  `json:"themeId"`
	NameKey          string               `json:"nameKey"`
	NameKeySecondary string               `json:"nameKeySecondary"`
	Schedule         []TournamentPhaseDto `json:"schedule"`
}

type TournamentPhaseDto struct {
	Id               int  `json:"id"`
	RegistrationTime int  `json:"registrationTime"`
	StartTime        int  `json:"startTime"`
	Cancelled        bool `json:"cancelled"`
}

type ClashService struct {
	client *RiotApi
}

// returns a list of active Clash players for a given summoner ID.
// If a summoner registers for multiple tournaments at the same time (e.g., Saturday and Sunday)
// then both registrations would appear in this list
//
// https://developer.riotgames.com/apis#clash-v1/GET_getPlayersBySummoner
func (s ClashService) PlayersBySummonerId(summonerId string) ([]PlayerDto, error) {
	endpoint := fmt.Sprintf("%s/lol/clash/v1/players/by-summoner/%s", s.client.PlatformUrl, summonerId)
	req, err := s.client.newRequest("GET", endpoint, nil)
	if err != nil {
		return []PlayerDto{}, err
	}
	body, err := s.client.do(req)
	if err != nil {
		return []PlayerDto{}, err
	}
	var players []PlayerDto
	err = json.Unmarshal(body, &players)
	if err != nil {
		return []PlayerDto{}, err
	}
	return players, nil
}

// get team by Id
//
// https://developer.riotgames.com/apis#clash-v1/GET_getTeamById
func (s ClashService) TeamByTeamId(teamId string) (TeamDto, error) {
	endpoint := fmt.Sprintf("%s/lol/clash/v1/teams/%s", s.client.PlatformUrl, teamId)
	req, err := s.client.newRequest("GET", endpoint, nil)
	if err != nil {
		return TeamDto{}, err
	}
	body, err := s.client.do(req)
	if err != nil {
		return TeamDto{}, err
	}
	var team TeamDto
	err = json.Unmarshal(body, &team)
	if err != nil {
		return TeamDto{}, err
	}
	return team, nil
}

// Get all active or upcoming tournaments
//
// https://developer.riotgames.com/apis#clash-v1/GET_getTournaments
func (s ClashService) GetTournaments() ([]TournamentDto, error) {
	endpoint := fmt.Sprintf("%s/lol/clash/v1/tournaments", s.client.PlatformUrl)
	req, err := s.client.newRequest("GET", endpoint, nil)
	if err != nil {
		return []TournamentDto{}, err
	}
	body, err := s.client.do(req)
	if err != nil {
		return []TournamentDto{}, err
	}
	var tournies []TournamentDto
	err = json.Unmarshal(body, &tournies)
	if err != nil {
		return []TournamentDto{}, err
	}
	return tournies, nil
}

// Get tournament by team Id
//
// https://developer.riotgames.com/apis#clash-v1/GET_getTournamentByTeam
func (s ClashService) GetTournamentByTeamId(teamId string) (TournamentDto, error) {
	endpoint := fmt.Sprintf("%s/lol/clash/v1/tournaments/by-team/%s", s.client.PlatformUrl, teamId)
	req, err := s.client.newRequest("GET", endpoint, nil)
	if err != nil {
		return TournamentDto{}, err
	}
	body, err := s.client.do(req)
	if err != nil {
		return TournamentDto{}, err
	}
	var tourny TournamentDto
	err = json.Unmarshal(body, &tourny)
	if err != nil {
		return TournamentDto{}, err
	}
	return tourny, nil
}

// Get tournament by tournament Id
//
// https://developer.riotgames.com/apis#clash-v1/GET_getTournamentById
func (s ClashService) GetTournamentById(tournamentId string) (TournamentDto, error) {
	endpoint := fmt.Sprintf("%s/lol/clash/v1/tournaments/%s", s.client.PlatformUrl, tournamentId)
	req, err := s.client.newRequest("GET", endpoint, nil)
	if err != nil {
		return TournamentDto{}, err
	}
	body, err := s.client.do(req)
	if err != nil {
		return TournamentDto{}, err
	}
	var tourny TournamentDto
	err = json.Unmarshal(body, &tourny)
	if err != nil {
		return TournamentDto{}, err
	}
	return tourny, nil
}
