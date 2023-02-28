package models

type TeamYearlyRecord struct {
	Year   int
	Wins   int
	Losses int
	Ties   int
}

func (t *TeamYearlyRecord) Init(year int) {
	t.Year = year
	t.Wins = 0
	t.Losses = 0
	t.Ties = 0
}