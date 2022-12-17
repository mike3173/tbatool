package models

// Custom type to hold value for eventType
type EventType int

// Declare related constants for each event type
const (
	Regional              EventType = iota // EnumIndex = 0
	District                               // EnumIndex = 1
	District_Cmp                           // EnumIndex = 2
	Cmp_Division                           // EnumIndex = 3
	Cmp_Finals                             // EnumIndex = 4
	District_Cmp_Division                  // EnumIndex = 5
	FOC                                    // EnumIndex = 6
	Remote                                 // EnumIndex = 7
	OffSeason             EventType = 99
	PreSeason             EventType = 100
	Unlabeled             EventType = -1
)

func (e EventType) String() string {
	if e == OffSeason {
		return "OffSeason"
	}
	if e == PreSeason {
		return "PreSeason"
	}
	if e == Unlabeled {
		return "--"
	}
	return [...]string{"Regional", "District", "District Championship Division", "District Championship", "Championship Division", "Championship Finals", "Festival of Champions", "Remote"}[e]
}
