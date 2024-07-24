package riotapi

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

//returns a list of active Clash players for a given summoner ID. 
//If a summoner registers for multiple tournaments at the same time (e.g., Saturday and Sunday) 
//then both registrations would appear in this list
//
//https://developer.riotgames.com/apis#clash-v1/GET_getPlayersBySummoner
func (s ClashService) PlayersBySummonerId(SummonerId string) ([]PlayerDto, error){
	return []PlayerDto{}, nil
}