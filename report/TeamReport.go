package report

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/mike3173/tbatool/models"
	"github.com/xuri/excelize/v2"
)

const WS_TEAMINFO = "team_info"
const WS_EVENT_PARTICIPATION = "event_participation"
const WS_HISTORICAL_MATCHES = "historical_matches"
const WS_HISTORICAL_DATA = "historical_data"
const WS_AWARDS = "awards"

func TeamReport(teamHistory models.TeamHistory) {
	// Create new xlsx file
	var fname = fmt.Sprintf("%s-historical-performace.xlsx", teamHistory.Team.Key)
	xlf := excelize.NewFile()

	xlf.SetSheetName("Sheet1", WS_TEAMINFO) // rename Sheet1
	xlf.NewSheet(WS_EVENT_PARTICIPATION)
	xlf.NewSheet(WS_HISTORICAL_MATCHES)
	xlf.NewSheet(WS_HISTORICAL_DATA)
	xlf.NewSheet(WS_AWARDS)

	sheetIndex := xlf.GetSheetIndex(WS_TEAMINFO)
	formatTeamInfoSheet(teamHistory, sheetIndex, xlf)

	sheetIndex = xlf.GetSheetIndex(WS_EVENT_PARTICIPATION)
	formatEventParticipationSheet(teamHistory, sheetIndex, xlf)

	sheetIndex = xlf.GetSheetIndex(WS_HISTORICAL_MATCHES)
	formatHistoricalMatchesSheet(teamHistory, sheetIndex, xlf)

	sheetIndex = xlf.GetSheetIndex(WS_HISTORICAL_DATA)
	formatHistoryDataSheet(teamHistory, sheetIndex, xlf)

	sheetIndex = xlf.GetSheetIndex(WS_AWARDS)
	formatAwardsSheet(teamHistory, sheetIndex, xlf)

	xlf.SetActiveSheet(xlf.GetSheetIndex(WS_TEAMINFO))
	// Save and close
	if err := xlf.SaveAs(fname); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func fixTeamKey(teamKey string) int {
	t := strings.TrimPrefix(teamKey, "frc")
	rtnValue, err := strconv.Atoi(t)
	if err != nil {
		panic(err)
	}
	return rtnValue
}

func formatAwardsSheet(teamHistory models.TeamHistory, sheetIdx int, xlf *excelize.File) {
	var awards []models.Award = teamHistory.Awards

	xlf.SetCellValue(WS_AWARDS, "A1", "YEAR")
	xlf.SetCellValue(WS_AWARDS, "B1", "EVENT")
	xlf.SetCellValue(WS_AWARDS, "C1", "AWARD")

}

func formatHistoryDataSheet(teamHistory models.TeamHistory, sheetIdx int, xlf *excelize.File) {
	var matches []models.Match = teamHistory.Matches
	histData := map[int]*models.TeamHistoryPerformance{}

	xlf.SetCellValue(WS_HISTORICAL_DATA, "A1", "TEAM #")
	xlf.SetCellValue(WS_HISTORICAL_DATA, "B1", "PARTNER PLAYS")
	xlf.SetCellValue(WS_HISTORICAL_DATA, "C1", "PARTNER WINS")
	xlf.SetCellValue(WS_HISTORICAL_DATA, "D1", "PARTNER LOSSES")
	xlf.SetCellValue(WS_HISTORICAL_DATA, "E1", "PARTNER TIES")
	xlf.SetCellValue(WS_HISTORICAL_DATA, "F1", "OPPONENT PLAYS")
	xlf.SetCellValue(WS_HISTORICAL_DATA, "G1", "OPPONENT WINS")
	xlf.SetCellValue(WS_HISTORICAL_DATA, "H1", "OPPONENT LOSSES")
	xlf.SetCellValue(WS_HISTORICAL_DATA, "I1", "OPPONENT TIES")
	xlf.SetCellValue(WS_HISTORICAL_DATA, "J1", "TOTAL PLAYED")
	xlf.SetCellValue(WS_HISTORICAL_DATA, "K1", "TOTAL WINS")
	xlf.SetCellValue(WS_HISTORICAL_DATA, "L1", "TOTAL LOSSES")
	xlf.SetCellValue(WS_HISTORICAL_DATA, "M1", "TOTAL TIES")
	xlf.SetCellValue(WS_HISTORICAL_DATA, "N1", "PARTNER WIN %")
	xlf.SetCellValue(WS_HISTORICAL_DATA, "O1", "OPPONENT WIN %")
	xlf.SetCellValue(WS_HISTORICAL_DATA, "P1", "TOTAL WIN %")
	xlf.SetCellValue(WS_HISTORICAL_DATA, "Q1", "POWER METRIC")
	xlf.SetCellValue(WS_HISTORICAL_DATA, "R1", "TBA LINK")

	// create historical data map
	var row int = 2 // data starts on row 2
	var partners models.MatchAllianceData
	var opponents models.MatchAllianceData

	fmt.Printf("Extracting historical data for Team %d\n", teamHistory.Team.TeamNumber)
	for i := 0; i < len(matches); i++ {
		fmt.Printf("Processing match %s %s, %d ", matches[i].EventKey, matches[i].GetCompLevelString(), matches[i].MatchNumber)
		if Contains(matches[i].Alliances.Red.TeamKeys, teamHistory.Team.Key) { // red alliance are partners, blue alliance are opponents
			partners = matches[i].Alliances.Red
			opponents = matches[i].Alliances.Blue

		} else { // blue alliance are partners, red alliance are opponents
			partners = matches[i].Alliances.Blue
			opponents = matches[i].Alliances.Red
		}
		// process partner list
		for p := 0; p < len(partners.TeamKeys); p++ {
			partnerTeam := fixTeamKey(partners.TeamKeys[p])
			if partnerTeam != teamHistory.Team.TeamNumber { // skip over the team whose history we are processing
				_, exists := histData[partnerTeam]
				if !exists {
					histData[partnerTeam] = &models.TeamHistoryPerformance{}
					histData[partnerTeam].Init(partnerTeam)
				}
				if partners.Score > opponents.Score { // partners win, opponents loss
					histData[partnerTeam].AddPartnerWin()
				} else if partners.Score < opponents.Score { // partners loss, opponents win
					histData[partnerTeam].AddPartnerLoss()
				} else { // tie
					histData[partnerTeam].AddPartnerTie()
				}
			}
		}
		// process opponent list
		for o := 0; o < len(opponents.TeamKeys); o++ {
			opponentTeam := fixTeamKey(opponents.TeamKeys[o])
			_, exists := histData[opponentTeam]
			if !exists {
				histData[opponentTeam] = &models.TeamHistoryPerformance{}
				histData[opponentTeam].Init(opponentTeam)
			}
			if partners.Score > opponents.Score { // partners win, opponents loss
				histData[opponentTeam].AddOpponentWin()
			} else if partners.Score < opponents.Score { // partners loss, opponents win
				histData[opponentTeam].AddOpponentLoss()
			} else { // tie
				histData[opponentTeam].AddOpponentTie()
			}
		}
	}
	// iterate over histData map, extract key slice
	keySlice := make([]int, 0)
	for key := range histData {
		keySlice = append(keySlice, key)
	}
	sort.Ints(keySlice) // sort the key slice
	for _, key := range keySlice {
		t := histData[key]

		xlf.SetCellValue(WS_HISTORICAL_DATA, fmt.Sprintf("A%d", row), t.TeamNumber)     // TEAM #
		xlf.SetCellValue(WS_HISTORICAL_DATA, fmt.Sprintf("B%d", row), t.PartnerPlays)   // PARTNER PLAYS
		xlf.SetCellValue(WS_HISTORICAL_DATA, fmt.Sprintf("C%d", row), t.PartnerWins)    // PARTNER WINS
		xlf.SetCellValue(WS_HISTORICAL_DATA, fmt.Sprintf("D%d", row), t.PartnerLosses)  // PARTNER LOSSES
		xlf.SetCellValue(WS_HISTORICAL_DATA, fmt.Sprintf("E%d", row), t.PartnerTies)    // PARTNER TIES
		xlf.SetCellValue(WS_HISTORICAL_DATA, fmt.Sprintf("F%d", row), t.OpponentPlays)  // OPPONENT PLAYS
		xlf.SetCellValue(WS_HISTORICAL_DATA, fmt.Sprintf("G%d", row), t.OpponentWins)   // OPPONENT WINS
		xlf.SetCellValue(WS_HISTORICAL_DATA, fmt.Sprintf("H%d", row), t.OpponentLosses) // OPPONENT LOSSES
		xlf.SetCellValue(WS_HISTORICAL_DATA, fmt.Sprintf("I%d", row), t.OpponentTies)   // OPPONENT TIES
		xlf.SetCellValue(WS_HISTORICAL_DATA, fmt.Sprintf("J%d", row), t.TotalPlays)    // TOTAL PLAYED
		xlf.SetCellValue(WS_HISTORICAL_DATA, fmt.Sprintf("K%d", row), t.TotalWins)      // TOTAL WINS
		xlf.SetCellValue(WS_HISTORICAL_DATA, fmt.Sprintf("L%d", row), t.TotalLosses)    // TOTAL LOSSES
		xlf.SetCellValue(WS_HISTORICAL_DATA, fmt.Sprintf("M%d", row), t.TotalTies)      // TOTAL TIES
		xlf.SetCellValue(WS_HISTORICAL_DATA, fmt.Sprintf("N%d", row), t.PartnerWinPct)  // PARTNER WIN %
		xlf.SetCellValue(WS_HISTORICAL_DATA, fmt.Sprintf("O%d", row), t.OpponentWinPct) // OPPONENT WIN %
		xlf.SetCellValue(WS_HISTORICAL_DATA, fmt.Sprintf("P%d", row), t.TotalWinPct)    // TOTAL WIN %
		xlf.SetCellValue(WS_HISTORICAL_DATA, fmt.Sprintf("Q%d", row), t.PowerMetric)    // POWER METRIC
		xlf.SetCellValue(WS_HISTORICAL_DATA, fmt.Sprintf("R%d", row), t.TBALink)        // TBA LINK
		row++
	}
}

func formatHistoricalMatchesSheet(teamHistory models.TeamHistory, sheetIdx int, xlf *excelize.File) {
	var matches []models.Match = teamHistory.Matches

	xlf.SetCellValue(WS_HISTORICAL_MATCHES, "A1", "EVENT_KEY")
	xlf.SetCellValue(WS_HISTORICAL_MATCHES, "B1", "MATCH TYPE")
	xlf.SetCellValue(WS_HISTORICAL_MATCHES, "C1", "MATCH NBR")
	xlf.SetCellValue(WS_HISTORICAL_MATCHES, "D1", "RED 1")
	xlf.SetCellValue(WS_HISTORICAL_MATCHES, "E1", "RED 2")
	xlf.SetCellValue(WS_HISTORICAL_MATCHES, "F1", "RED 3")
	xlf.SetCellValue(WS_HISTORICAL_MATCHES, "G1", "BLUE 1")
	xlf.SetCellValue(WS_HISTORICAL_MATCHES, "H1", "BLUE 2")
	xlf.SetCellValue(WS_HISTORICAL_MATCHES, "I1", "BLUE 3")
	xlf.SetCellValue(WS_HISTORICAL_MATCHES, "J1", "RED SCORE")
	xlf.SetCellValue(WS_HISTORICAL_MATCHES, "K1", "BLUE SCORE")
	xlf.SetCellValue(WS_HISTORICAL_MATCHES, "L1", "PARTNER 1")
	xlf.SetCellValue(WS_HISTORICAL_MATCHES, "M1", "PARTNER 2")
	xlf.SetCellValue(WS_HISTORICAL_MATCHES, "N1", "PARTNER 3")
	xlf.SetCellValue(WS_HISTORICAL_MATCHES, "O1", "ALLIANCE COLOR")
	xlf.SetCellValue(WS_HISTORICAL_MATCHES, "P1", "PART W/L/T")
	xlf.SetCellValue(WS_HISTORICAL_MATCHES, "Q1", "OPPONENT 1")
	xlf.SetCellValue(WS_HISTORICAL_MATCHES, "R1", "OPPONENT 2")
	xlf.SetCellValue(WS_HISTORICAL_MATCHES, "S1", "OPPONENT 3")
	xlf.SetCellValue(WS_HISTORICAL_MATCHES, "T1", "OPP W/L/T")
	xlf.SetCellValue(WS_HISTORICAL_MATCHES, "U1", "ALL LOOKUP KEY")
	var row int = 2 // data starts on row 2
	for i := 0; i < len(matches); i++ {
		xlf.SetCellValue(WS_HISTORICAL_MATCHES, fmt.Sprintf("A%d", row+i), matches[i].EventKey)
		xlf.SetCellValue(WS_HISTORICAL_MATCHES, fmt.Sprintf("B%d", row+i), fmt.Sprintf("%d-%s", matches[i].GetCompLevelKey(), matches[i].GetCompLevelString()))
		xlf.SetCellValue(WS_HISTORICAL_MATCHES, fmt.Sprintf("C%d", row+i), matches[i].MatchNumber)
		xlf.SetCellValue(WS_HISTORICAL_MATCHES, fmt.Sprintf("D%d", row+i), fixTeamKey(matches[i].Alliances.Red.TeamKeys[0]))
		xlf.SetCellValue(WS_HISTORICAL_MATCHES, fmt.Sprintf("E%d", row+i), fixTeamKey(matches[i].Alliances.Red.TeamKeys[1]))
		xlf.SetCellValue(WS_HISTORICAL_MATCHES, fmt.Sprintf("F%d", row+i), fixTeamKey(matches[i].Alliances.Red.TeamKeys[2]))
		xlf.SetCellValue(WS_HISTORICAL_MATCHES, fmt.Sprintf("G%d", row+i), fixTeamKey(matches[i].Alliances.Blue.TeamKeys[0]))
		xlf.SetCellValue(WS_HISTORICAL_MATCHES, fmt.Sprintf("H%d", row+i), fixTeamKey(matches[i].Alliances.Blue.TeamKeys[1]))
		xlf.SetCellValue(WS_HISTORICAL_MATCHES, fmt.Sprintf("I%d", row+i), fixTeamKey(matches[i].Alliances.Blue.TeamKeys[2]))
		xlf.SetCellValue(WS_HISTORICAL_MATCHES, fmt.Sprintf("J%d", row+i), matches[i].Alliances.Red.Score)
		xlf.SetCellValue(WS_HISTORICAL_MATCHES, fmt.Sprintf("K%d", row+i), matches[i].Alliances.Blue.Score)
		if Contains(matches[i].Alliances.Red.TeamKeys, teamHistory.Team.Key) { // team was red alliance
			xlf.SetCellValue(WS_HISTORICAL_MATCHES, fmt.Sprintf("L%d", row+i), fixTeamKey(matches[i].Alliances.Red.TeamKeys[0]))
			xlf.SetCellValue(WS_HISTORICAL_MATCHES, fmt.Sprintf("M%d", row+i), fixTeamKey(matches[i].Alliances.Red.TeamKeys[1]))
			xlf.SetCellValue(WS_HISTORICAL_MATCHES, fmt.Sprintf("N%d", row+i), fixTeamKey(matches[i].Alliances.Red.TeamKeys[2]))
			xlf.SetCellValue(WS_HISTORICAL_MATCHES, fmt.Sprintf("O%d", row+i), "RED")
			if matches[i].WinningAlliance == "red" {
				xlf.SetCellValue(WS_HISTORICAL_MATCHES, fmt.Sprintf("P%d", row+i), "W")
				xlf.SetCellValue(WS_HISTORICAL_MATCHES, fmt.Sprintf("T%d", row+i), "L")
			} else if matches[i].WinningAlliance == "blue" {
				xlf.SetCellValue(WS_HISTORICAL_MATCHES, fmt.Sprintf("P%d", row+i), "L")
				xlf.SetCellValue(WS_HISTORICAL_MATCHES, fmt.Sprintf("T%d", row+i), "W")
			} else {
				xlf.SetCellValue(WS_HISTORICAL_MATCHES, fmt.Sprintf("P%d", row+i), "T")
				xlf.SetCellValue(WS_HISTORICAL_MATCHES, fmt.Sprintf("T%d", row+i), "T")
			}
			xlf.SetCellValue(WS_HISTORICAL_MATCHES, fmt.Sprintf("Q%d", row+i), fixTeamKey(matches[i].Alliances.Blue.TeamKeys[0]))
			xlf.SetCellValue(WS_HISTORICAL_MATCHES, fmt.Sprintf("R%d", row+i), fixTeamKey(matches[i].Alliances.Blue.TeamKeys[1]))
			xlf.SetCellValue(WS_HISTORICAL_MATCHES, fmt.Sprintf("S%d", row+i), fixTeamKey(matches[i].Alliances.Blue.TeamKeys[2]))
		} else { // team was blue alliance
			xlf.SetCellValue(WS_HISTORICAL_MATCHES, fmt.Sprintf("L%d", row+i), fixTeamKey(matches[i].Alliances.Blue.TeamKeys[0]))
			xlf.SetCellValue(WS_HISTORICAL_MATCHES, fmt.Sprintf("M%d", row+i), fixTeamKey(matches[i].Alliances.Blue.TeamKeys[1]))
			xlf.SetCellValue(WS_HISTORICAL_MATCHES, fmt.Sprintf("N%d", row+i), fixTeamKey(matches[i].Alliances.Blue.TeamKeys[2]))
			xlf.SetCellValue(WS_HISTORICAL_MATCHES, fmt.Sprintf("O%d", row+i), "BLUE")
			if matches[i].WinningAlliance == "red" {
				xlf.SetCellValue(WS_HISTORICAL_MATCHES, fmt.Sprintf("P%d", row+i), "L")
				xlf.SetCellValue(WS_HISTORICAL_MATCHES, fmt.Sprintf("T%d", row+i), "W")
			} else if matches[i].WinningAlliance == "blue" {
				xlf.SetCellValue(WS_HISTORICAL_MATCHES, fmt.Sprintf("P%d", row+i), "W")
				xlf.SetCellValue(WS_HISTORICAL_MATCHES, fmt.Sprintf("T%d", row+i), "L")
			} else {
				xlf.SetCellValue(WS_HISTORICAL_MATCHES, fmt.Sprintf("P%d", row+i), "T")
				xlf.SetCellValue(WS_HISTORICAL_MATCHES, fmt.Sprintf("T%d", row+i), "T")
			}
			xlf.SetCellValue(WS_HISTORICAL_MATCHES, fmt.Sprintf("Q%d", row+i), fixTeamKey(matches[i].Alliances.Red.TeamKeys[0]))
			xlf.SetCellValue(WS_HISTORICAL_MATCHES, fmt.Sprintf("R%d", row+i), fixTeamKey(matches[i].Alliances.Red.TeamKeys[1]))
			xlf.SetCellValue(WS_HISTORICAL_MATCHES, fmt.Sprintf("S%d", row+i), fixTeamKey(matches[i].Alliances.Red.TeamKeys[2]))
		}
		xlf.SetCellValue(WS_HISTORICAL_MATCHES, fmt.Sprintf("U%d", row+i), matches[i].Alliances.GetAllianceLookupKey())
	}
}

func formatEventParticipationSheet(teamHistory models.TeamHistory, sheetIdx int, xlf *excelize.File) {
	var events []models.Event = teamHistory.Events

	// Set Header Cell values
	xlf.SetCellValue(WS_EVENT_PARTICIPATION, "A1", "EVENT_KEY")
	xlf.SetCellValue(WS_EVENT_PARTICIPATION, "B1", "NAME")
	xlf.SetCellValue(WS_EVENT_PARTICIPATION, "C1", "EVENT_TYPE")
	xlf.SetCellValue(WS_EVENT_PARTICIPATION, "D1", "LOCATION_NAME")
	xlf.SetCellValue(WS_EVENT_PARTICIPATION, "E1", "ADDRESS")
	xlf.SetCellValue(WS_EVENT_PARTICIPATION, "F1", "START_DATE")
	xlf.SetCellValue(WS_EVENT_PARTICIPATION, "G1", "END_DATE")

	var row int = 2 // data starts on row 2
	for i := 0; i < len(events); i++ {
		xlf.SetCellValue(WS_EVENT_PARTICIPATION, fmt.Sprintf("A%d", row+i), events[i].Key)
		xlf.SetCellValue(WS_EVENT_PARTICIPATION, fmt.Sprintf("B%d", row+i), events[i].Name)
		xlf.SetCellValue(WS_EVENT_PARTICIPATION, fmt.Sprintf("C%d", row+i), events[i].EventType)
		xlf.SetCellValue(WS_EVENT_PARTICIPATION, fmt.Sprintf("D%d", row+i), events[i].LocationName)
		xlf.SetCellValue(WS_EVENT_PARTICIPATION, fmt.Sprintf("E%d", row+i), events[i].Address)
		xlf.SetCellValue(WS_EVENT_PARTICIPATION, fmt.Sprintf("F%d", row+i), events[i].StartDate)
		xlf.SetCellValue(WS_EVENT_PARTICIPATION, fmt.Sprintf("G%d", row+i), events[i].EndDate)
	}
}

func formatTeamInfoSheet(teamHistory models.TeamHistory, sheetIdx int, xlf *excelize.File) {
	var team models.Team = teamHistory.Team

	// Set Cell values
	xlf.SetCellValue(WS_TEAMINFO, "A1", "Team Info")
	xlf.SetCellValue(WS_TEAMINFO, "A2", "FRC Team:")
	xlf.SetCellValue(WS_TEAMINFO, "B2", fmt.Sprintf("%d %s", team.TeamNumber, team.Nickname))
	xlf.SetCellValue(WS_TEAMINFO, "A3", "Rookie Year:")
	xlf.SetCellValue(WS_TEAMINFO, "B3", fmt.Sprintf("%d", team.RookieYear))
	xlf.SetCellValue(WS_TEAMINFO, "A4", "From:")
	xlf.SetCellValue(WS_TEAMINFO, "B4", fmt.Sprintf("%s %s, %s %s", team.SchoolName, team.City, team.StateProv, team.Country))

	// Format Sheet1, named
	if err := xlf.MergeCell(WS_TEAMINFO, "A1", "B1"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Create default Style
	defStyle, err := xlf.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Size: 18},
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defStyleBold, err := xlf.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold: true,
			Size: 18},
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defStyleBoldCenter, err := xlf.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold: true,
			Size: 18},
		Alignment: &excelize.Alignment{
			Horizontal: "center"},
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := xlf.SetCellStyle(WS_TEAMINFO, "A1", "B4", defStyle); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := xlf.SetCellStyle(WS_TEAMINFO, "A1", "A2", defStyleBoldCenter); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := xlf.SetCellStyle(WS_TEAMINFO, "A2", "A4", defStyleBold); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := xlf.SetColWidth(WS_TEAMINFO, "A", "A", 19); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := xlf.SetColWidth(WS_TEAMINFO, "B", "B", 78); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func Contains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}
