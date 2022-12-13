package models

type Match struct {
	Key              string
	Comp_level       string
	Set_number       int
	Match_number     int
	Alliances        []MatchAlliance
	Winning_alliance string
	Event_key        string
	Time             int64
	Actual_time      int64
	Predicted_time   int64
	Post_result_time int64
	Score_breakdown  string
	Videos           []Videos
}
