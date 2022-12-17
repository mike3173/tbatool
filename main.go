package main

import (
	"encoding/json"
	"fmt"

	"github.com/mike3173/tbatool/models"
	"github.com/mike3173/tbatool/services"
)

func main() {
	fmt.Println("Calling API...")
	var bodyBytes []byte = services.GetTeamInfo("frc3173")

	// var t models.Team
	// json.Unmarshal(bodyBytes, &t)
	// fmt.Printf("team: %+v\n", t)

	// bodyBytes = services.GetTeamYearsParticipated("frc3173")
	// var years_participated []int
	// json.Unmarshal(bodyBytes, &years_participated)
	// fmt.Printf("years participated: %+v len=%d\n", years_participated, len(years_participated))

	bodyBytes = services.GetTeamEventsForYear("frc3173", 2022)
	var events []models.Event
	json.Unmarshal(bodyBytes, &events)
	// fmt.Printf("events: %+v\n", events)

	// bodyBytes = services.GetMatch("2022nyro", "qm", 43)
	// var match []models.Match
	// json.Unmarshal(bodyBytes, &match)
	// fmt.Printf("match: %+v\n", match)

	// report.TeamReport(t)

	for e := models.Regional; e <= models.Remote; e++ {
		fmt.Print(e)
		fmt.Printf(" eventtype=%d (%s)\n", e, e.String())
	}

	for p := models.Bracket8Team; p <= models.Custom; p++ {
		fmt.Print(p)
		fmt.Printf(" playofftype=%d (%s)\n", p, p.String())
	}

	fmt.Println("report complete")
}
