package report

import (
	"fmt"

	"github.com/mike3173/tbatool/models"
)

func MatchReportCsv(match models.Match) {
	fmt.Printf("%+v\n", match.Key)
	var alliance string = "blue"
	fmt.Printf("    %s teams %d\n", alliance, len(match.Alliances.Blue["team_keys"]))
	// for i := 0; i < len(match.Alliances.Blue["team_keys"].(string)); i+ {
	// 	fmt.Printf("%+v, %d, %s", match.Alliances.Blue["team_keys"][i].(string), i, alliane)
	// }
	fmt.Printf("    %s teams %+v\n", alliance, match.Alliances.Blue["team_keys"])
	fmt.Printf("    %s surrogate %+v\n", alliance, match.Alliances.Blue["surroate_team_keys"])
	fmt.Printf("    %s dq %+v\n", alliance, match.Alliances.Blue["dq_team_key"])
	fmt.Printf("    %s score %+v\n", alliance, match.Alliances.Blue["score"])

	alliance = "red"
	fmt.Printf("    %s teams %+v\n", alliance, match.Alliances.Red["team_keys"])
	if match.Alliances.Red["surrogate_team_keys"] != nil {
		fmt.Printf("    %s surrogate %+v\n", alliance, match.Alliances.Red["surrogate_team_keys"])
	}
	fmt.Printf("    %s dq %+v\n", alliance, match.Alliances.Red["dq_team_key"])
	fmt.Printf("    %s score %+v\n", alliance, match.Alliances.Red["score"])
}
