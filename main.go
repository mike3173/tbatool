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

	var t models.Team
	json.Unmarshal(bodyBytes, &t)
	fmt.Printf("team: %+v\n", t)

	bodyBytes = services.GetTeamYearsParticipated("frc3173")
	var years_participated []int
	json.Unmarshal(bodyBytes, &years_participated)
	fmt.Printf("years participated: %+v len=%d\n", years_participated, len(years_participated))

	bodyBytes = services.GetTeamEvents("frc3173")
	var events []models.Event
	json.Unmarshal(bodyBytes, &events)
	fmt.Printf("events: %+v\n", events)
}
