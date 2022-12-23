package models

import "fmt"

type MatchScoreBreakdown2022 struct {
	AdjustPoints            int    `json:"adjustPoints"`
	AutoCargoLowerBlue      int    `json:"autoCargoLowerBlue"`
	AutoCargoLowerFar       int    `json:"autoCargoLowerFar"`
	AutoCargoLowerNear      int    `json:"autoCargoLowerNear"`
	AutoCargoLowerRed       int    `json:"autoCargoLowerRed"`
	AutoCargoPoints         int    `json:"autoCargoPoints"`
	AutoCargoTotal          int    `json:"autoCargoTotal"`
	AutoCargoUpperBlue      int    `json:"autoCargoUpperBlue"`
	AutoCargoUpperFar       int    `json:"autoCargoUpperFar"`
	AutoCargoUpperNear      int    `json:"autoCargoUpperNear"`
	AutoCargoUpperRed       int    `json:"autoCargoUpperRed"`
	AutoPoints              int    `json:"autoPoints"`
	AutoTaxiPoints          int    `json:"autoTaxiPoints"`
	CargoBonusRankingPoint  bool   `json:"cargoBonusRankingPoint"`
	EndgamePoints           int    `json:"endgamePoints"`
	EndgameRobot1           string `json:"endgameRobot1"`
	EndgameRobot2           string `json:"endgameRobot2"`
	EndgameRobot3           string `json:"endgameRobot3"`
	FoulCount               int    `json:"foulCount"`
	FoulPoints              int    `json:"foulPoints"`
	HangarBonusRankingPoint bool   `json:"hangarBonusRankingPoint"`
	MatchCargoTotal         int    `json:"matchCargoTotal"`
	QuintetAchieved         bool   `json:"quintetAchieved"`
	Rp                      int    `json:"rp"`
	TaxiRobot1              string `json:"taxiRobot1"`
	TaxiRobot2              string `json:"taxiRobot2"`
	TaxiRobot3              string `json:"taxiRobot3"`
	TechFoulCount           int    `json:"techFoulCount"`
	TeleopCargoLowerBlue    int    `json:"teleopCargoLowerBlue"`
	TeleopCargoLowerFar     int    `json:"teleopCargoLowerFar"`
	TeleopCargoLowerNear    int    `json:"teleopCargoLowerNear"`
	TeleopCargoLowerRed     int    `json:"teleopCargoLowerRed"`
	TeleopCargoPoints       int    `json:"teleopCargoPoints"`
	TeleopCargoTotal        int    `json:"teleopCargoTotal"`
	TeleopCargoUpperBlue    int    `json:"teleopCargoUpperBlue"`
	TeleopCargoUpperFar     int    `json:"teleopCargoUpperFar"`
	TeleopCargoUpperNear    int    `json:"teleopCargoUpperNear"`
	TeleopCargoUpperRed     int    `json:"teleopCargoUpperRed"`
	TeleopPoints            int    `json:"teleopPoints"`
	TotalPoints             int    `json:"totalPoints"`
}

func (m MatchScoreBreakdown2022) getRedScoreData(teamAllianceNbr int, alliance string) string {
	return fmt.Sprintf("%d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %t, %d, %s, %d, %d, %t, %d, %t, %d, %s, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d",
		m.AdjustPoints, m.AutoCargoLowerFar, m.AutoCargoLowerNear, m.AutoCargoLowerRed, m.AutoCargoPoints, m.AutoCargoTotal, m.AutoCargoUpperFar, m.AutoCargoUpperNear, m.AutoCargoUpperRed, m.AutoPoints, m.AutoTaxiPoints, m.CargoBonusRankingPoint, m.EndgamePoints, m.getEndgameRobotValue(teamAllianceNbr), m.FoulCount, m.FoulPoints, m.HangarBonusRankingPoint, m.MatchCargoTotal, m.QuintetAchieved, m.Rp, m.getTaxiRobotValue(teamAllianceNbr), m.TechFoulCount, m.TeleopCargoLowerFar, m.TeleopCargoLowerNear, m.TeleopCargoLowerRed, m.TeleopCargoPoints, m.TeleopCargoTotal, m.TeleopCargoUpperFar, m.TeleopCargoUpperNear, m.TeleopCargoUpperRed, m.TeleopPoints, m.TotalPoints)
}

func (m MatchScoreBreakdown2022) getBlueScoreData(teamAllianceNbr int, alliance string) string {
	return fmt.Sprintf("%d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %t, %d, %s, %d, %d, %t, %d, %t, %d, %s, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d",
		m.AdjustPoints, m.AutoCargoLowerBlue, m.AutoCargoLowerFar, m.AutoCargoLowerNear, m.AutoCargoPoints, m.AutoCargoTotal, m.AutoCargoUpperBlue, m.AutoCargoUpperFar, m.AutoCargoUpperNear, m.AutoPoints, m.AutoTaxiPoints, m.CargoBonusRankingPoint, m.EndgamePoints, m.getEndgameRobotValue(teamAllianceNbr), m.FoulCount, m.FoulPoints, m.HangarBonusRankingPoint, m.MatchCargoTotal, m.QuintetAchieved, m.Rp, m.getTaxiRobotValue(teamAllianceNbr), m.TechFoulCount, m.TeleopCargoLowerBlue, m.TeleopCargoLowerFar, m.TeleopCargoLowerNear, m.TeleopCargoPoints, m.TeleopCargoTotal, m.TeleopCargoUpperBlue, m.TeleopCargoUpperFar, m.TeleopCargoUpperNear, m.TeleopPoints, m.TotalPoints)
}

func (m MatchScoreBreakdown2022) GetScoreData(teamAllianceNbr int, alliance string) string {
	var rtnValue string = ""

	if alliance == "blue" {
		rtnValue = m.getBlueScoreData(teamAllianceNbr, alliance)
	} else if alliance == "red" {
		rtnValue = m.getRedScoreData(teamAllianceNbr, alliance)
	} else {
		rtnValue = fmt.Sprintf("unknown alliance %s", alliance)
	}
	return rtnValue
}

func (m MatchScoreBreakdown2022) getTaxiRobotValue(teamAllianceNbr int) string {
	var result string

	switch teamAllianceNbr {
	case 1:
		result = m.TaxiRobot1
	case 2:
		result = m.TaxiRobot2
	case 3:
		result = m.TaxiRobot3
	default:
		result = "unknown"
	}
	return result
}

func (m MatchScoreBreakdown2022) getEndgameRobotValue(teamAllianceNbr int) string {
	var result string

	switch teamAllianceNbr {
	case 1:
		result = m.EndgameRobot1
	case 2:
		result = m.EndgameRobot2
	case 3:
		result = m.EndgameRobot3
	default:
		result = "unknown"
	}
	return result
}