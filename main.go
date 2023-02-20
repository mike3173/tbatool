package main

import (
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
			teamArg = "frc" +teamArg
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
			teamArg = "frc" +teamArg
		}
		fmt.Println("Getting Team History for " + strings.TrimPrefix(teamArg, "frc"))
		reportLine = fmt.Sprintf("Team History for %s", strings.TrimPrefix(teamArg, "frc"))
		getTeamHistory(teamArg)
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
			th.Events = append(th.Events, events[ii])
		}
	}

	for i := 0; i < len(th.Events); i++ {
		if th.Events[i].EventType >= 0 && th.Events[i].EventType <= 7 {		// skip offseason type events
			bodyBytes = services.GetTeamEventMatches(team, th.Events[i].Key)
			var matches []models.Match
			json.Unmarshal(bodyBytes, &matches)
			fmt.Printf("Getting matches for event %s\n", th.Events[i].Key)
			for ii := 0; ii < len(matches); ii++ {
				th.Matches = append(th.Matches, matches[ii])
			}
		}
	}

	bodyBytes = services.GetTeamAwards(team)
	json.Unmarshal(bodyBytes, &th.Awards)

	fmt.Println("Writing report ...")
	report.TeamReport(th)
}

func getTeamAwards(team string) {
	var bodyBytes = services.GetTeamAwards(team)
	var ta []models.Award
	json.Unmarshal(bodyBytes, &ta)

	// fmt.Printf("%+v\n", string(bodyBytes))
	fmt.Printf("%+v\n", ta)
}