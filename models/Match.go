package models

type Match struct {
	Key             string              `json:"key"`
	CompLevel       string              `json:"comp_level"`
	SetNumber       int                 `json:"set_number"`
	MatchNumber     int                 `json:"match_number"`
	Alliances       MatchAlliance       `json:"alliances"`
	WinningAlliance string              `json:"winning_alliance"`
	EventKey        string              `json:"event_key"`
	Time            int64               `json:"time"`
	ActualTime      int64               `json:"actual_time"`
	PredictedTime   int64               `json:"predicted_time"`
	PostResultTime  int64               `json:"post_result_time"`
	ScoreBreakdown  MatchScoreBreakdown `json:"score_breakdown"`
	Videos          []Videos            `json:"videos"`
}
