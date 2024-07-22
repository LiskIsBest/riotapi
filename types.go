package riotapi

type Status struct {
	Message     string `json:"message"`
	Status_code string `json:"status_code"`
}

type ResError struct {
	Status Status `json:"status"`
}

type AccountDto struct {
	Puuid    string `json:"puuid"`
	GameName string `json:"gameName"`
	TagLine  string `json:"tagLine"`
}
