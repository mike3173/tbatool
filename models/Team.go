package models

type Team struct {
	Key               string            // TBA team key with the format frcXXXX with XXXX representing the team number.
	Team_number       int               // Official team number issued by FIRST.
	Nickname          string            // Team nickname provided by FIRST.
	Name              string            // Official long name registered with FIRST.
	School_name       string            // Name of team school or affilited group registered with FIRST.
	City              string            // City of team derived from parsing the address registered with FIRST.
	State_prov        string            // State of team derived from parsing the address registered with FIRST.
	Country           string            // Country of team derived from parsing the address registered with FIRST.
	Address           string            // Will be NULL, for future development.
	Postal_code       string            // Postal code from the team address.
	Gmaps_place_id    string            // Will be NULL, for future development.
	Gmaps_url         string            // Will be NULL, for future development.
	Lat               float32           // Will be NULL, for future development.
	Lng               float32           // Will be NULL, for future development.
	Location_name     string            // Will be NULL, for future development.
	Website           string            // Official website associated with the team.
	Rookie_year       int               // First year the team officially competed.
	Motto             string            // Team's motto as provided by FIRST. This field is deprecated and will return null - will be removed at end-of-season in 2019.
	Home_championship map[string]string // Location of the team's home championship each year as a key-value pair. The year (as a string) is the key, and the city is the value.
}
