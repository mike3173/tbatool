package models

type Event struct {
	Key                 string
	Name                string
	Event_code          string
	Event_type          int
	District            District_List
	City                string
	State_prov          string
	Country             string
	Start_date          string
	End_date            string
	Year                int
	Short_name          string
	Event_type_string   string
	Week                int
	Address             string
	Postal_code         string
	Gmaps_place_id      string
	Gmaps_url           string
	Lat                 float32
	Lng                 float32
	Location_name       string
	Timezone            string
	Website             string
	First_event_id      string
	First_event_code    string
	Webcasts            []Webcast
	Division_keys       []string
	Parent_event_key    string
	Playoff_type        int
	Playoff_type_string string
}
