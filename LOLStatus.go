package riotapi

import "errors"

type ContentDto struct {
	Locale  string `json:"locale"`
	Content string `json:"content"`
}

type UpdateDto struct {
	Id                int          `json:"id"`
	Author            string       `json:"author"`
	Publish           bool         `json:"publish"`
	Publish_locations []string     `json:"publish_locations"`
	Translations      []ContentDto `json:"translations"`
	Created_at        string       `json:"created_at"`
	Updated_at        string       `json:"updated_at"`
}

type StatusDto struct {
	Id                 int          `json:"id"`
	Maintenance_status string       `json:"maintenance_status"` // Legal values: scheduled, in_progress, complete
	Incident_severity  string       `json:"incident_severity"`  // Legal values: info, warning, critical
	Titles             []ContentDto `json:"titles"`
	Updates            []UpdateDto  `json:"updates"`
	Created_at         string       `json:"created_at"`
	Archive_at         string       `json:"archive_at"`
	Updated_at         string       `json:"updated_at"`
	Platforms          []string     `json:"platforms"` //Legal values: windows, macos, android, ios, ps4, xbone, switch
}

type PlatformDataDto struct {
	Id           string      `json:"id"`
	Name         string      `json:"name"`
	Locales      []string    `json:"locales"`
	Maintenances []StatusDto `json:"maintenances"`
	Incidents    []StatusDto `json:"incidents"`
}

type LOLStatusService struct {
	client *RiotApi
}

// Get League of Legends status for the given platform
//
// https://developer.riotgames.com/apis#lol-status-v4/GET_getPlatformData
func (s LOLStatusService) PlatformData() (PlatformDataDto, error) {
	return PlatformDataDto{}, errors.New("LOLStatus.PlatformData not implemented yet.")
}
