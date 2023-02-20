package models

type TeamHistory struct {
	Team              Team
	YearsParticipated []int
	Events            []Event
	Matches           []Match
	Awards            []Award
}