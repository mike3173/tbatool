package services

import "fmt"

const GETMATCHURL = "https://www.thebluealliance.com/api/v3/match/%s"

func GetMatch(eventKey string, matchType string, matchNumber int) []byte {
	var url string = fmt.Sprintf(GETMATCHURL, fmt.Sprintf("%s_%s%d", eventKey, matchType, matchNumber))
	var bodyBytes []byte = TBAMakeRequest(url)
	return bodyBytes
}
