package report

import (
	"fmt"
	"os"
	"strings"

	"github.com/mike3173/tbatool/models"
)

func isInList(element string, list []string) bool {
	var result bool = false

	if len(list) > 0 {
		for i := 0; i < len(list); i++ {
			if list[i] == element {
				result = true
				break
			}
		}
	}
	return result
}

func MatchReportCsv(fileName string, match models.Match) int {
	var dataLine string
	var alliance string = "blue"
	var lines int = 0

	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(match.Alliances.Blue.TeamKeys); i++ {
		var teamAllianceNbr int = i + 1
		dataLine = fmt.Sprintf("%s,%s,%d,%d,%d,", match.EventKey, match.CompLevel, match.GetCompLevelKey(), match.SetNumber, match.MatchNumber)
		dataLine += fmt.Sprintf("%s,%d,%s,", alliance, teamAllianceNbr, strings.Replace(match.Alliances.Blue.TeamKeys[i], "frc", "", 1))
		dataLine += fmt.Sprintf("%+v,", isInList(match.Alliances.Blue.TeamKeys[i], match.Alliances.Blue.SurrogateTeamKeys))
		dataLine += fmt.Sprintf("%+v,", isInList(match.Alliances.Blue.TeamKeys[i], match.Alliances.Blue.DqTeamKeys))
		dataLine += match.ScoreBreakdown.Blue.GetScoreData(teamAllianceNbr, alliance)
		dataLine += "\n"
		_, err := f.WriteString(dataLine)
		if err != nil {
			panic(err)
		}
		lines++
	}

	alliance = "red"
	for i := 0; i < len(match.Alliances.Red.TeamKeys); i++ {
		var teamAllianceNbr int = i + 1
		dataLine = fmt.Sprintf("%s,%s,%d,%d,%d,", match.EventKey, match.CompLevel, match.GetCompLevelKey(), match.SetNumber, match.MatchNumber)
		dataLine += fmt.Sprintf("%s,%d,%s,", alliance, teamAllianceNbr, strings.Replace(match.Alliances.Red.TeamKeys[i], "frc", "", 1))
		dataLine += fmt.Sprintf("%+v,", isInList(match.Alliances.Red.TeamKeys[i], match.Alliances.Red.SurrogateTeamKeys))
		dataLine += fmt.Sprintf("%+v,", isInList(match.Alliances.Red.TeamKeys[i], match.Alliances.Red.DqTeamKeys))
		dataLine += match.ScoreBreakdown.Red.GetScoreData(teamAllianceNbr, alliance)
		dataLine += "\n"
		_, err := f.WriteString(dataLine)
		if err != nil {
			panic(err)
		}
		lines++
	}
	f.Sync()
	f.Close()
	return lines
}

func GetMatchHeaderLine() string {
	return "event_key,comp_level,comp_level_key,set_number,match_number,alliance,team_alliance_number,team_number,is_surrogate,is_disqualified,adjustPoints,autoCargoLower,autoCargoLowerFar,autoCargoLowerNear,autoCargoPoints,autoCargoTotal,autoCargoUpper,autoCargoUpperFar,autoCargoUpperNear,autoPoints,autoTaxiPoints,cargoBonusRankingPoint,endgamePoints,endgameRobot,foulCount,foulPoints,hangarBonusRankingPoint,matchCargoTotal,quintetAchieved,rp,taxiRobot,techFoulCount,teleopCargoLower,teleopCargoLowerFar,teleopCargoLowerNear,teleopCargoPoints,teleopCargoTotal,teleopCargoUpper,teleopCargoUpperFar,teleopCargoUpperNear,teleopPoints,totalPoints\n"
}
