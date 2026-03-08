package riotapi

import "errors"

// need to make enums for State, Tracking, and Level

type ChallengeConfigInfoDto struct {
	Id             int                          `json:"id"`
	LocalizedNames map[string]map[string]string `json:"localizedNames"`
	State          string                       `json:"state"`    // need to figure out if this is the right data type when api works again
	Tracking       string                       `json:"tracking"` // need to figure out if this is the right data type when api works again
	StartTimestamp int                          `json:"startTimestamp"`
	EndTimestamp   int                          `json:"endTimestamp"`
	Leaderboard    bool                         `json:"leaderboard"`
	Threshholds    map[string]float64
}

type ChallengeInfoDto struct {
	Percentile float64 `json:"percentile"`
	PlayersInLevel int `json:"playersInLevel"`
	AchievedTime int `json:"achievedTime"`
	Value float64 `json:"value"`
	ChallengeId int `json:"challengeId"`
	Level string `json:"level"` // Legal values: NONE, IRON, BRONZE, SILVER, GOLD, PLATINUM, DIAMOND, MASTER, GRANDMASTER, CHALLENGER, HIGHEST_NOT_LEADERBOARD_ONLY, HIGHEST, LOWEST
	Position int `json:"position"`
}

type PlayerClientPreferencesDto struct {
	BannerAccent string `json:"bannerAccent"`
	Title string `json:"title"`
	ChallengeIds []string `json:"ChallengeIds"`
	CrestBorder string `json:"crestBorder"`
	PrestigeCrestBorderLevel int `json:"prestigeCrestBorderLevel"`
}

type ChallengePointDto struct{
	Level string `json:"level"`
	Current int `json:"current"`
	Max int `json:"max"`
	Precentile int `json:"Precentile"` // weird spelling error in api documentation. need to verify when api working
}

type PlayerInfoDto struct {
	Challenges []ChallengeInfoDto `json:"challenges"`
	Preferences PlayerClientPreferencesDto `json:"preferences"`
	TotalPoints ChallengePointDto `json:"totalPoints"`
	CategoryPoints map[string]ChallengePointDto `json:"categoryPoints"`
}

type ApexPlayerInfoDto struct {
	Puuid    string  `json:"puuid"`
	Value    float64 `json:"value"`
	Position int     `json:"position"`
}

type LOLChallengesService struct {
	client *RiotApi
}

// List of all basic challenge configuration information (includes all translations for names and descriptions)
//
// https://developer.riotgames.com/apis#lol-challenges-v1/GET_getAllChallengeConfigs
func (s LOLChallengesService) ChallengesConfig() ([]ChallengeConfigInfoDto, error) {
	return []ChallengeConfigInfoDto{}, errors.New("LOLChallengesService.ChallengesConfig not implemented yet.")
}

// Map of level to percentile of players who have achieved it - keys: ChallengeId -> Season -> Level -> percentile of players who achieved it
//
// CHECK WHAT "LEVEL" IS WHEN API IS WORKING
//
// https://developer.riotgames.com/apis#lol-challenges-v1/GET_getAllChallengePercentiles
func (s LOLChallengesService) ChallengesPercent() (map[int]map[int]map[string]float64, error) {
	return map[int]map[int]map[string]float64{}, errors.New("LOLChallengesService.ChallengesPercent not implemented yet.")
}

// Get challenge configuration
//
// https://developer.riotgames.com/apis#lol-challenges-v1/GET_getChallengeConfigs
func (s LOLChallengesService) ChallengesByChallengeId() (ChallengeConfigInfoDto, error) {
	return ChallengeConfigInfoDto{}, errors.New("LOLChallengesService.ChallengesByChallengeId not implemented yet.")
}

// Return top players for each level. Level must be MASTER, GRANDMASTER or CHALLENGER.
//
// https://developer.riotgames.com/apis#lol-challenges-v1/GET_getChallengeLeaderboards
func (s LOLChallengesService) ChallengesByChallengeIdByLevel() ([]ApexPlayerInfoDto, error) {
	return []ApexPlayerInfoDto{}, errors.New("LOLChallengesService.ChallengesByChallengeIdByLevel not implemented yet.")
}

// Map of level to percentile of players who have achieved it
//
// CHECK WHAT "LEVEL" IS WHEN API IS WORKING
//
// https://developer.riotgames.com/apis#lol-challenges-v1/GET_getChallengePercentiles
func (s LOLChallengesService) ChallengesByChallengeIdPercent() (map[string]float64, error) {
	return map[string]float64{}, errors.New("LOLChallengesService.ChallengesByChallengeIdPercent not implemented yet.")
}

// Returns player information with list of all progressed challenges
//
// https://developer.riotgames.com/apis#lol-challenges-v1/GET_getPlayerData
func (s LOLChallengesService) PlayerDataByPuuid(puuid string) (PlayerInfoDto, error){
	return PlayerInfoDto{}, errors.New("LOLChallengesService.PlayerDataByPuuid not implemented yet.")
}
