package models

type TeamHistory struct {
	Team              Team
	YearsParticipated []int
	Events            []Event
	Matches           []Match
	Awards            []Award
	YearlyRecords     map[int]TeamYearlyRecord
}

func (t *TeamHistory) Init(teamNbr string) {
	t.YearlyRecords = make(map[int]TeamYearlyRecord)
}