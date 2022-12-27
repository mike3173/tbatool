package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/mike3173/tbatool/models"
	"github.com/mike3173/tbatool/report"
	"github.com/mike3173/tbatool/services"
)

func main() {
	fmt.Println("Calling API...")

	// bodyBytes = services.GetTeamInfo("frc3173")
	// var t models.Team
	// json.Unmarshal(bodyBytes, &t)
	// fmt.Printf("team: %+v\n\n", t)

	// bodyBytes = services.GetTeamYearsParticipated("frc3173")
	// var yearsParticipated []int
	// json.Unmarshal(bodyBytes, &yearsParticipated)
	// fmt.Printf("years participated: %+v len=%d\n\n", yearsParticipated, len(yearsParticipated))

	// bodyBytes = services.GetTeamEventsForYear("frc3173", 2022)
	// var events []models.Event
	// json.Unmarshal(bodyBytes, &events)
	// fmt.Printf("events: %+v\n\n", events)

	var bodyBytes []byte = services.GetEventMatches("2022ohcl")
	var matches []models.Match
	err := json.Unmarshal(bodyBytes, &matches)
	if err != nil {
		panic(err)
	} else {
		var headerLine bool = false
		var outFileName string
		var lines int = 0

		for i := 0; i < len(matches); i++ {
			if !headerLine {
				outFileName = matches[i].EventKey + "_match_data.csv"
				f, err := os.Create(outFileName)
				if err != nil {
					panic(err)
				}

				n, err := f.WriteString(report.GetMatchHeaderLine())
				if err != nil {
					panic(err)
				}
				fmt.Printf("wrote header (%d bytes)\n", n)
				headerLine = true
				f.Sync()
				f.Close()
			}
			lines += report.MatchReportCsv(outFileName, matches[i])
		}
		fmt.Printf("wrote %d lines of data\n", lines)
	}

	// bodyBytes []byte = services.GetMatch("2022nyro", "qm", 43)
	// var match models.Match
	// err := json.Unmarshal(bodyBytes, &match)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// } else {
	// 	fmt.Printf("match.key: %s\n", match.Key)
	// 	fmt.Printf("match.Alliances: BLUE %+v\n", match.Alliances.Blue)
	// 	fmt.Printf("match.Alliances: RED  %+v\n", match.Alliances.Red)
	// 	fmt.Printf("match.scorebreakdown: BLUE %+v\n", match.ScoreBreakdown.Blue)
	// 	fmt.Printf("match.scorebreakdown: RED  %+v\n", match.ScoreBreakdown.Red)
	// 	sortedKeys := make([]string, 0, len(match.ScoreBreakdown.Blue))
	// 	fmt.Printf("\nmatch score key list:\n")
	// 	for key := range match.ScoreBreakdown.Blue {
	// 		sortedKeys = append(sortedKeys, key)
	// 	}
	// 	sort.Strings(sortedKeys)
	// 	fmt.Printf("sorted key=%+v\n", sortedKeys)
	// 	var header string
	// 	for key := range sortedKeys {
	// 		if len(header) == 0 {
	// 			header = fmt.Sprintf("%+v", sortedKeys[key])
	// 		} else {
	// 			header = fmt.Sprintf("%s, %+v", header, sortedKeys[key])
	// 		}
	// 	}
	// 	fmt.Printf("HEADER: %s\n", header)
	// 	fmt.Println()
	// 	fmt.Printf("match score value list:\n")
	// 	var dataRow string
	// 	for _, key := range sortedKeys {
	// 		if len(dataRow) == 0 {
	// 			dataRow = fmt.Sprintf("%+v", match.ScoreBreakdown.Blue[key])
	// 		} else {
	// 			dataRow = fmt.Sprintf("%s, %+v", dataRow, match.ScoreBreakdown.Blue[key])
	// 		}
	// 	}
	// 	fmt.Printf("DATAROW: %s\n", dataRow)
	// 	fmt.Println()

	// var blueMap map[string]interface{}
	// var redMap map[string]interface{}
	// dataBlue, _ := json.Marshal(match.ScoreBreakdown.Blue)
	// json.Unmarshal(dataBlue, &blueMap)
	// dataRed, _ := json.Marshal(match.ScoreBreakdown.Red)
	// json.Unmarshal(dataRed, &redMap)
	// fmt.Printf("match.scorebreakdown: BLUE %+v\n", blueMap)
	// fmt.Printf("match.scorebreakdown: RED  %+v\n", redMap)
	// sortedKeys := make([]string, 0, len(blueMap))
	// for key := range blueMap {
	// 	sortedKeys = append(sortedKeys, key)
	// 	fmt.Printf("%s ", key)
	// }
	// sort.Strings(sortedKeys)
	// fmt.Println()
	// for _, key := range sortedKeys {
	// 	fmt.Printf(" %+v", blueMap[key])
	// }
	// fmt.Println()
	// }

	// report.TeamReport(t)

	// for e := models.Regional; e <= models.Remote; e++ {
	// 	fmt.Print(e)
	// 	fmt.Printf(" eventtype=%d (%s)\n", e, e.String())
	// }

	// for p := models.Bracket8Team; p <= models.Custom; p++ {
	// 	fmt.Print(p)
	// 	fmt.Printf(" playofftype=%d (%s)\n", p, p.String())
	// }

	fmt.Println("report complete")
}
