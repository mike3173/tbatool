package report

import (
	"fmt"

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

func MatchReportCsv(match models.Match) {
	var dataLine string

	fmt.Printf("%+v\n", match.Key)
	var alliance string = "blue"
	fmt.Printf("    %s teams %+v\n", alliance, match.Alliances.Blue.TeamKeys)
	for i := 0; i < len(match.Alliances.Blue.TeamKeys); i++ {
		dataLine = fmt.Sprintf("%s, %s, %d, %d, ", match.EventKey, match.CompLevel, match.Setumber, match.MatchNumber)
		dataLine = fmt.Sprintf("%s, %s, %d, %s", dataLine, i+1, match.Alliances.Blue.TeamKeys[i], alliance)
		dataLine = fmt.Sprintf("%s, %+v\n", dataLine, isInList(match.Alliances.Blue.TeamKeys[i], match.Alliances.Blue.SurrogateTeamKeys))
	}
	fmt.Printf("    %s surrogate %+v\n", alliance, match.Alliances.Blue.SurrogateTeamKeys)
	fmt.Printf("    s dq %+v\n", alliance, match.Alliances.Blue.DqTeamKeys)
	fmt.Printf("    %s score %+v\n", alliance, match.Alliances.Blue.Score)

	alliance = "red"
	fmt.Printf("    %s teams %+v\n", alliance, match.Alliances.Red.TeamKeys)
	for i := 0; i < len(match.Alliances.Red.TeamKeys); i++ {
		fmt.Printf("%s, %d, %s\n", match.Alliances.Red.TeamKeys[i], i, alliance)
	}
	fmt.Printf("    %s surrogate %+v\n", alliance, match.Alliances.Red.SurrogateTeamKeys)
	fmt.Printf("    %s dq %+v\n", alliance, match.Alliances.Red.DqTeamKeys)
	fmt.Printf("    %s score %+v\n", alliance, match.Alliances.Red.Score)

}
