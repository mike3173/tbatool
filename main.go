package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/mike3173/tbatool/models"
)

func main() {
	fmt.Println("Calling API...")
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://www.thebluealliance.com/api/v3/team/frc3173", nil)
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
	var t models.Team
	json.Unmarshal(bodyBytes, &t)
	fmt.Printf("team: %+v\n", t)
	// var responseObject http.Response
	// json.Unmarshal(bodyBytes, &responseObject)
	// fmt.Printf("API Response as struct %+v\n", responseObject)

}
