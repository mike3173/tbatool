package services

import (
	"fmt"
)

const GETTEAMURL string = "https://www.thebluealliance.com/api/v3/team/%s"
const GETTEAMYEARSURL string = "https://www.thebluealliance.com/api/v3/team/%s/years_participated"
const GETTEAMEVENTSURL string = "https://www.thebluealliance.com/api/v3/team/%s/events"
const GETTEAMEVENTSSIMPLEURL string = "https://www.thebluealliance.com/api/v3/team/%s/events/simple"
const GETTEAMEVENTSYEARURL string = "https://www.thebluealliance.com/api/v3/team/%s/events/%d"
const GETTEAMEVENTSYEARSIMPLEURL string = "https://www.thebluealliance.com/api/v3/team/%s/events/%d/simple"

func GetTeamInfo(teamKey string) []byte {
	var url string = fmt.Sprintf(GETTEAMURL, teamKey)
	var bodyBytes []byte = TBAMakeRequest(url)
	return bodyBytes
}

func GetTeamYearsParticipated(teamKey string) []byte {
	var url string = fmt.Sprintf(GETTEAMYEARSURL, teamKey)
	var bodyBytes []byte = TBAMakeRequest(url)
	return bodyBytes
}

func GetTeamEvents(teamKey string) []byte {
	var url string = fmt.Sprintf(GETTEAMEVENTSURL, teamKey)
	var bodyBytes []byte = TBAMakeRequest(url)
	return bodyBytes
}

func GetTeamEventsSimple(teamKey string) []byte {
	var url string = fmt.Sprintf(GETTEAMEVENTSSIMPLEURL, teamKey)
	var bodyBytes []byte = TBAMakeRequest(url)
	return bodyBytes
}

func GetTeamEventsForYear(teamKey string, year int) []byte {
	var url string = fmt.Sprintf(GETTEAMEVENTSYEARURL, teamKey, year)
	var bodyBytes []byte = TBAMakeRequest(url)
	return bodyBytes
}

func GetTeamEventsForYearSimple(teamKey string, year int) []byte {
	var url string = fmt.Sprintf(GETTEAMEVENTSYEARSIMPLEURL, teamKey, year)
	var bodyBytes []byte = TBAMakeRequest(url)
	return bodyBytes
}
