package services

import "fmt"

const GETMATCHURL = "https://www.thebluealliance.com/api/v3/match/%s_%s%d"
const GETEVENTMATCHESURL = "https://www.thebluealliance.com/api/v3/event/%s/matches"
const GETEVENTMATCHESSIMPLEURL = "https://www.thebluealliance.com/api/v3/event/%s/matches/simple"

func GetMatch(eventKey string, matchType string, matchNumber int) []byte {
	var url string = fmt.Sprintf(GETMATCHURL, eventKey, matchType, matchNumber)
	var bodyBytes []byte = TBAMakeRequest(url)
	return bodyBytes
}

func GetEventMatches(eventKey string) []byte {
	var url string = fmt.Sprintf(GETEVENTMATCHESURL, eventKey)
	var bodyBytes []byte = TBAMakeRequest(url)
	return bodyBytes
}

func GetEventMatchesSimple(eventKey string) []byte {
	var url string = fmt.Sprintf(GETEVENTMATCHESSIMPLEURL, eventKey)
	var bodyBytes []byte = TBAMakeRequest(url)
	return bodyBytes
}
