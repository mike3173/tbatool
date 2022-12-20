package models

type MatchScoreBreakdown struct {
	Blue map[string]interface{} `json:"blue"`
	Red  map[string]interface{} `json:"red"`
	// Blue MatchScoreBreakdown2022 `json:"blue"`
	// Red  MatchScoreBreakdown2022 `json:"red"`
}
