package riotapi

import (
	"encoding/json"
	"fmt"
)

type ChampionMasteryDto struct {
	Puuid                        string                  `json:"puuid"`
	ChampionPointsUntilNextLevel int                     `json:"championPointsUntilNextLevel"`
	ChestGranted                 bool                    `json:"chestGranted"`
	ChampionId                   int                     `json:"championId"`
	LastPlayTime                 int                     `json:"lastPlayTime"`
	ChampionLevel                int                     `json:"championLevel"`
	ChampionPoints               int                     `json:"championPoints"`
	ChampionPointsSinceLastLevel int                     `json:"championPointsSinceLastLevel"`
	MarkRequiredForNextLevel     int                     `json:"markRequiredForNextLevel"`
	ChampionSeasonMilestone      int                     `json:"championSeasonMilestone"`
	NextSeasonMilestone          NextSeasonMilestonesDto `json:"nextSeasonMilestone"`
	TokensEarned                 int                     `json:"tokensEarned"`
	MilestoneGrades              []string                `json:"milestoneGrades"`
}

type NextSeasonMilestonesDto struct {
	RequiredGradeCounts RequiredGradeCounts `json:"requiredGradecounts"`
	RewardMarks         int                 `json:"rewardMarks"`
	Bonus               bool                `json:"bonus"`
	RewardConfig        RewardConfigDto     `json:"rewardConfig"`
}

type RewardConfigDto struct {
	RewardValue   string `json:"rewardValue"`
	RewardType    string `json:"rewardType"`
	MaximumReward int    `json:"maximumReward"`
}

type RequiredGradeCounts struct {
	S_plus  int `json:"S+"`
	S_norm  int `json:"S"`
	S_minus int `json:"S-"`
	A_plus  int `json:"A+"`
	A_norm  int `json:"A"`
	A_minus int `json:"A-"`
	B_plus  int `json:"B+"`
	B_norm  int `json:"B"`
	B_minus int `json:"B-"`
	C_plus  int `json:"C+"`
	C_norm  int `json:"C"`
	C_minus int `json:"C-"`
	D_plus  int `json:"D+"`
	D_norm  int `json:"D"`
	D_minus int `json:"D-"`
}

type ChampionMasteryService struct {
	client *RiotApi
}

func (s ChampionMasteryService) ChampionMasteries(puuid string) ([]ChampionMasteryDto, error) {
	endpoint := fmt.Sprintf("%s/lol/champion-mastery/v4/champion-masteries/by-puuid/%s",baseHostURL1,puuid)
	req, err := s.client.NewRequest("GET", endpoint, nil)
	if err != nil{
		return []ChampionMasteryDto{}, err
	}
	body, err := s.client.Do(req)
	if err != nil{
		return []ChampionMasteryDto{}, err
	}
	var masteries []ChampionMasteryDto
	err = json.Unmarshal(body, &masteries)
	if err != nil{
		return []ChampionMasteryDto{}, err
	}
	return masteries, nil
}
