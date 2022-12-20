package models

type Team struct {
	Key              string            `json:"key"`
	TeamNumber       int               `json:"team_number"`
	Nickname         string            `json:"nickname"`
	Name             string            `json:"name"`
	SchoolName       string            `json:"school_name"`
	City             string            `json:"city"`
	StateProv        string            `json:"state_prov"`
	Country          string            `json:"country"`
	Address          string            `json:"address"`
	PostalCode       string            `json:"postal_code"`
	GmapsPlaceId     string            `json:"gmaps_place_id"`
	GmapsUrl         string            `json:"gmaps_url"`
	Lat              float32           `json:"lat"`
	Lng              float32           `json:"lng"`
	LocationName     string            `json:"location_name"`
	Website          string            `json:"website"`
	RookieYear       int               `json:"rookie_year"`
	Motto            string            `json:"motto"`
	HomeChampionship map[string]string `json:"home_championship"`
}
