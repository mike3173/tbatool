package models

type MatchScoreBreakdown struct {
	Blue map[string]interface{} `json:"blue"`
	Red  map[string]interface{} `json:"red"`
}