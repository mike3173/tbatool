package models

import "fmt"

type TeamHistoryPerformance struct {
	TeamNumber     int
	PartnerPlays   int
	PartnerWins    int
	PartnerLosses  int
	PartnerTies    int
	OpponentPlays  int
	OpponentWins   int
	OpponentLosses int
	OpponentTies   int
	TotalWins      int
	TotalLosses    int
	TotalTies      int
	TotalPlays    int
	PartnerWinPct  float32
	OpponentWinPct float32
	TotalWinPct    float32
	PowerMetric    float32
	TBALink        string
}

func (t *TeamHistoryPerformance) Init(teamNbr int) {
	t.TeamNumber = teamNbr
	t.TBALink = fmt.Sprintf("https://www.thebluealliance.com/team/%d", teamNbr)

	t.PartnerPlays = 0
	t.PartnerWins = 0
	t.PartnerLosses = 0
	t.PartnerTies = 0
	t.OpponentPlays = 0
	t.OpponentWins = 0
	t.OpponentLosses = 0
	t.OpponentTies = 0
	t.TotalWins = 0
	t.TotalLosses = 0
	t.TotalTies = 0
	t.TotalPlays = 0
	t.PartnerWinPct = 0.0
	t.OpponentWinPct = 0.0
	t.TotalWinPct = 0.0
	t.PowerMetric = 0.0
}

func (t *TeamHistoryPerformance) AddPartnerWin() {
	t.PartnerPlays++
	t.PartnerWins++
	t.TotalPlays++
	t.TotalWins++

	t.PartnerWinPct = float32(t.PartnerWins) / float32(t.PartnerPlays)
	t.TotalWinPct = float32(t.TotalWins) / float32(t.TotalPlays)
	t.PowerMetric = t.PartnerWinPct - t.OpponentWinPct
}

func (t *TeamHistoryPerformance) AddPartnerLoss() {
	t.PartnerPlays++
	t.PartnerLosses++
	t.TotalPlays++
	t.TotalLosses++

	t.PartnerWinPct = float32(t.PartnerWins) / float32(t.PartnerPlays)
	t.TotalWinPct = float32(t.TotalWins) / float32(t.TotalPlays)
	t.PowerMetric = t.PartnerWinPct - t.OpponentWinPct
}

func (t *TeamHistoryPerformance) AddPartnerTie() {
	t.PartnerPlays++
	t.PartnerTies++
	t.TotalPlays++
	t.TotalTies++

	t.PartnerWinPct = float32(t.PartnerWins) / float32(t.PartnerPlays)
	t.TotalWinPct = float32(t.TotalWins) / float32(t.TotalPlays)
	t.PowerMetric = t.PartnerWinPct - t.OpponentWinPct
}

func (t *TeamHistoryPerformance) AddOpponentWin() {
	t.OpponentPlays++
	t.OpponentWins++
	t.TotalPlays++
	t.TotalWins++	

	t.OpponentWinPct = float32(t.OpponentWins) / float32(t.OpponentPlays)
	t.TotalWinPct = float32(t.TotalWins) / float32(t.TotalPlays)
	t.PowerMetric = t.PartnerWinPct - t.OpponentWinPct
}

func (t *TeamHistoryPerformance) AddOpponentLoss() {
	t.OpponentPlays++
	t.OpponentLosses++
	t.TotalPlays++
	t.TotalLosses++

	t.OpponentWinPct = float32(t.OpponentWins) / float32(t.OpponentPlays)
	t.TotalWinPct = float32(t.TotalWins) / float32(t.TotalPlays)
	t.PowerMetric = t.PartnerWinPct - t.OpponentWinPct
}

func (t *TeamHistoryPerformance) AddOpponentTie() {
	t.OpponentPlays++
	t.OpponentTies++
	t.TotalTies++

	t.OpponentWinPct = float32(t.OpponentWins) / float32(t.OpponentPlays)
	t.TotalWinPct = float32(t.TotalWins) / float32(t.TotalPlays)
	t.PowerMetric = t.PartnerWinPct - t.OpponentWinPct
}
