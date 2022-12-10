package services

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

const GETTEAMURL string = "https://www.thebluealliance.com/api/v3/team/%s"
const GETTEAMYEARSURL string = "https://www.thebluealliance.com/api/v3/team/%s/years_participated"
const GETTEAMEVENTSURL string = "https://www.thebluealliance.com/api/v3/team/%s/events"

func GetTeamInfo(teamKey string) []byte {
	var url string = fmt.Sprintf(GETTEAMURL, teamKey)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Print(err.Error())
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("X-TBA-Auth-Key", "hJm8rF1L3SiiqMDGEXc5LIwWvhCQBHyNDdJ9WaaiV8DuvnL2URPyMYkqGQoxlYhh")
	// fmt.Printf("req: %+v\n", req)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	// fmt.Printf("resp: %+v\n", resp)
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	fmt.Printf("bodybytes: %+v\n", string(bodyBytes))
	if err != nil {
		fmt.Print(err.Error())
	}
	return bodyBytes
}

func GetTeamYearsParticipated(teamKey string) []byte {
	var url string = fmt.Sprintf(GETTEAMYEARSURL, teamKey)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Print(err.Error())
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("X-TBA-Auth-Key", "hJm8rF1L3SiiqMDGEXc5LIwWvhCQBHyNDdJ9WaaiV8DuvnL2URPyMYkqGQoxlYhh")
	// fmt.Printf("req: %+v\n", req)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	// fmt.Printf("resp: %+v\n", resp)
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	fmt.Printf("bodybytes: %+v\n", string(bodyBytes))
	if err != nil {
		fmt.Print(err.Error())
	}
	return bodyBytes
}