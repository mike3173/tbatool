package models

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
