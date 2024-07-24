package riotapi

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

/*
link for getting a key: https://developer.riotgames.com/
enpoints: https://developer.riotgames.com/apis
riot docs:
	- https://developer.riotgames.com/docs/portal#_getting-started
	- https://developer.riotgames.com/docs/lol
*/

var baseUrl string = ".api.riotgames.com"

type RiotApi struct {
	client *http.Client

	Platform    string
	Region      string
	PlatformUrl string
	RegionalUrl string

	key            string
	Accounts       *AccountsService
	ChampMasteries *ChampionMasteryService
	Champions      *ChampionService
}

/*
Creates an instances of the RiotApi struct.

Valid region options: NA1, BR1, LA1, LA2, KR, JP1, EUN1, EUW1,
TR1, RU, OC1, PG2, SG2, TH2, TW2, VN2, NA, BR, LAN, LAS, JP,
EUNE, EUW, TR, OCE, PG, SG, TH, TW, VN,
*/
func NewRiotApi(apiKey string, region string) (*RiotApi, error) {
	r := &RiotApi{}
	client := http.Client{}

	regional, platform, err := convertPlatform2Regional(region)
	if err != nil{
		return &RiotApi{}, err
	}

	r.Platform = platform
	r.Region = region
	r.PlatformUrl = fmt.Sprintf("https://%s%s",platform,baseUrl)
	r.RegionalUrl = fmt.Sprintf("https://%s%s",regional,baseUrl)

	r.client = &client
	r.key = apiKey
	r.Accounts = &AccountsService{client: r}
	r.ChampMasteries = &ChampionMasteryService{client: r}
	r.Champions = &ChampionService{client: r}
	return r, nil
}

func (r RiotApi) NewRequest(method string, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-Riot-Token", r.key)
	return req, nil
}

func (r RiotApi) Do(req *http.Request) ([]byte, error) {
	resp, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func convertPlatform2Regional(region string) (string, string, error) {
    var regional string

    switch strings.ToUpper(region) {
    case "NA1", "BR1", "LA1", "LA2", "NA", "BR", "LAN", "LAS":
        regional = "AMERICAS"
    case "KR", "JP1", "JP":
        regional = "ASIA"
    case "EUN1", "EUW1", "TR1", "RU", "EUNE", "EUW", "TR":
        regional = "EUROPE"
    case "OC1", "PH2", "SG2", "TH2", "TW2", "VN2", "OCE", "PH", "SG", "TH", "TW", "VN":
        regional = "SEA"
    default:
        return "", "", fmt.Errorf("entered region: %s, is not a valid region", region)
    }
	return regional, region, nil
}