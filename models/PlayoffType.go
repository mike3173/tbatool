package models

// Custom type to hold value for playoffType
type PlayoffType int

// Declare related constants for each playoff type
const (
	Bracket8Team PlayoffType = iota
	Bracket16Team
	Bracket4Team
	AvgScore8Team
	RoundRobin6Team
	DoubleElim8Team
	BestOf5Finals
	BestOf3Finals
	Custom
)

func (p PlayoffType) String() string {
	return [...]string{"Elimination Bracket (8 Alliances)", "Elimination Bracket (4 Alliances)", "Elimination Bracket (16 Alliances)", "Average Score (8 Alliances)", "Round Robin (6 Alliances)", "Double Elimination Bracket (8 Alliances)", "Best of 3 Finals", "Best of 5 Finals", "Custom"}[p]
}
