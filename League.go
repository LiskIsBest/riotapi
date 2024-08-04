package riotapi

type LeagueListDto struct {
	LeagueId string          `json:"leagueId"`
	Entries  []LeagueItemDto `json:"entries"`
	Tier     string          `json:"tier"`
	Name     string          `json:"name"`
	Queue    string          `json:"queue"`
}

type LeagueItemDto struct {
	
}

type MiniSeriesDto struct {
}

type LeagueService struct {
	client *RiotApi
}
