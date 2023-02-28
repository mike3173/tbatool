package models

import (
	"fmt"
	"strings"
)

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

func (ma MatchAlliance) GetAllianceLookupKey() string {
	var rtnValue string = ""
	for i := 0; i < len(ma.Red.TeamKeys); i++ {
		rtnValue += fmt.Sprintf(":%s", strings.TrimPrefix(ma.Red.TeamKeys[i], "frc"))
	}
	for i := 0; i < len(ma.Blue.TeamKeys); i++ {
		rtnValue += fmt.Sprintf(":%s", strings.TrimPrefix(ma.Blue.TeamKeys[i], "frc"))
	}
	rtnValue += ":"
	return rtnValue
}

func (ma MatchAlliance) IsTeamRedAlliance(team string) bool {
	for _, v := range ma.Red.TeamKeys {
		if v == team {
			return true
		}
	}
	return false
}

func (ma MatchAlliance) IsTeamBlueAlliance(team string) bool {
	for _, v := range ma.Blue.TeamKeys {
		if v == team {
			return true
		}
	}
	return false
}