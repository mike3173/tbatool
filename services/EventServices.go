package services

import "fmt"

const GETEVENTSYEAR = "https://www.thebluealliance.com/api/v3/events/%d"
const GETEVENTSYEARSIMPLE = "https://www.thebluealliance.com/api/v3/events/%d/simple"

func GetEventsYear(year int) []byte {
	var url string = fmt.Sprintf(GETEVENTSYEAR, year)
	var bodyBytes []byte = TBAMakeRequest(url)
	return bodyBytes
}

func GetEventsYearSimple(year int) []byte {
	var url string = fmt.Sprintf(GETEVENTSYEARSIMPLE, year)
	var bodyBytes []byte = TBAMakeRequest(url)
	return bodyBytes
}
