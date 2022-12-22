package models

type MatchAlliance struct {
	Blue MatchAllianceData `json:"blue"`
	Red  MatchAllianceData `json:"red"`
}

type MatchAllianceData struct {
	TeamKeys          []string `json:"team_keys"`
	DqTeamKeys        []string `json:"dq_team_keys"`
	SurrogateTeamKeys []string `json:"surrogate_team_keys"`
	Score             int      `json:"score"`
}
