package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/mike3173/tbatool/models"
	"github.com/mike3173/tbatool/report"
	"github.com/mike3173/tbatool/services"
)

func main() {
	if len(os.Args) == 1 {
		usage("help")
		os.Exit(0)
	}
	cmdArg := os.Args[1]
	reportLine := ""

	switch cmdArg {
	case "team:awards":
		teamArg := os.Args[2]
		if !strings.HasPrefix(teamArg, "frc") {
			teamArg = "frc" + teamArg
		}
		fmt.Println("Getting Awards for Team " + strings.TrimPrefix(teamArg, "frc"))
		reportLine = fmt.Sprintf("Award History for Team %s", strings.TrimPrefix(teamArg, "frc"))
		getTeamAwards(teamArg)
	case "event:matches":
		eventArg := os.Args[2]
		fmt.Println("Getting Matches for Event " + eventArg)
		reportLine = fmt.Sprintf("Event Matches for %s", eventArg)
		getEventMatches(eventArg)
	case "team:history":
		teamArg := os.Args[2]
		if !strings.HasPrefix(teamArg, "frc") {
			teamArg = "frc" + teamArg
		}
		fmt.Println("Getting Team History for " + strings.TrimPrefix(teamArg, "frc"))
		reportLine = fmt.Sprintf("Team History for %s", strings.TrimPrefix(teamArg, "frc"))
		getTeamHistory(teamArg)
	case "list:team:history":
		listArg := os.Args[2]
		fileReader, err := os.Open(listArg)
		if err != nil {
			panic(err)
		}
		scanner := bufio.NewScanner(fileReader)
		for scanner.Scan() {
			teamArg := scanner.Text()
			fmt.Println("Getting Team History for " + strings.TrimPrefix(teamArg, "frc"))
			getTeamHistory(teamArg)
		}
		reportLine = fmt.Sprintf("Team History from file %s", listArg)
	case "help":
		usage("help")
	default:
		usage("unknown command")
	}
	fmt.Printf("%s complete\n", reportLine)
}

func usage(description string) {
	fmt.Printf("%s\n\n", description)
	fmt.Printf("\thelp - This text\n")
	fmt.Printf("\tteam:history team_number - History report for an FRC team eg. frc3173\n")
	fmt.Printf("\tevent:matches event_key - Match data for an FRC event eg. 2022nyro\n")
}

func getEventMatches(event string) {
	var bodyBytes []byte = services.GetEventMatches(event)
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
}

func getTeamHistory(team string) {
	var bodyBytes = services.GetTeamInfo(team)
	var th models.TeamHistory
	th.Init(team)
	json.Unmarshal(bodyBytes, &th.Team)
	fmt.Println("Getting years participated ...")

	bodyBytes = services.GetTeamYearsParticipated(team)
	json.Unmarshal(bodyBytes, &th.YearsParticipated)

	for i := 0; i < len(th.YearsParticipated); i++ {
		bodyBytes = services.GetTeamEventsForYear(team, th.YearsParticipated[i])
		var events []models.Event
		json.Unmarshal(bodyBytes, &events)
		fmt.Printf("Getting events for %d\n", th.YearsParticipated[i])
		for ii := 0; ii < len(events); ii++ {
			if reportEvent(events[ii]) {
				fmt.Printf("adding event %s eventType=%s\n", events[ii].Key, events[ii].EventType.String())
				th.Events = append(th.Events, events[ii])
			} else {
				fmt.Printf("skipping event %s eventType=%s\n", events[ii].Key, events[ii].EventType.String())
			}
		}
	}

	for i := 0; i < len(th.Events); i++ {
		theEvent := th.Events[i]
		bodyBytes = services.GetTeamEventMatches(team, theEvent.Key)
		var matches []models.Match
		json.Unmarshal(bodyBytes, &matches)
		fmt.Printf("Getting matches for event %s eventType=%s\n", th.Events[i].Key, th.Events[i].EventType.String())
		for ii := 0; ii < len(matches); ii++ {
			m := matches[ii]
			th.Matches = append(th.Matches, m)

			yr, ok := th.YearlyRecords[theEvent.Year]
			if !ok {
				yr.Init(theEvent.Year)
			}
			if m.Alliances.IsTeamBlueAlliance(team) {
				if m.Alliances.Blue.Score > m.Alliances.Red.Score {
					yr.Wins++
				} else if m.Alliances.Blue.Score < m.Alliances.Red.Score {
					yr.Losses++
				} else {
					yr.Ties++
				}
			} else if m.Alliances.IsTeamRedAlliance(team) {
				if m.Alliances.Red.Score > m.Alliances.Blue.Score {
					yr.Wins++
				} else if m.Alliances.Red.Score < m.Alliances.Blue.Score {
					yr.Losses++
				} else {
					yr.Ties++
				}
			}
			th.YearlyRecords[theEvent.Year] = yr
		}
	}

	bodyBytes = services.GetTeamAwards(team)
	json.Unmarshal(bodyBytes, &th.Awards)

	fmt.Println("Writing report ...")
	report.TeamReportExcel(th)
}

func getTeamAwards(team string) {
	var bodyBytes = services.GetTeamAwards(team)
	var ta []models.Award
	json.Unmarshal(bodyBytes, &ta)

	// fmt.Printf("%+v\n", string(bodyBytes))
	fmt.Printf("%+v\n", ta)
}

func reportEvent(event models.Event) bool {
	var rtnValue bool

	if event.EventType.InSeasonEvent() && event.Year >= 2005 { // skip offseason type events and dDon't report legacy matches before 3v3 model
		rtnValue = true
	}
	return rtnValue
}
