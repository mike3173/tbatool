package models

// Custom type to hold value for eventType
type EventType int

// Declare related constants for each event type
const (
	Regional            EventType = 0
	District            EventType = 1
	DistrictCmp         EventType = 2
	CmpDivision         EventType = 3
	CmpFinals           EventType = 4
	DistrictCmpDivision EventType = 5
	FOC                 EventType = 6
	Remote              EventType = 7
	OffSeason           EventType = 99
	PreSeason           EventType = 100
	Unlabeled           EventType = -1
)

func (e EventType) InSeasonEvent() bool {
	rtnValue := false
	switch e {
	case Regional:
		rtnValue = true
	case District:
		rtnValue = true
	case DistrictCmp:
		rtnValue = true
	case CmpDivision:
		rtnValue = true
	case CmpFinals:
		rtnValue = true
	case DistrictCmpDivision:
		rtnValue = true
	case FOC:
		rtnValue = true
	case Remote:
		rtnValue = true
	case OffSeason:
		rtnValue = false
	case PreSeason:
		rtnValue = false
	case Unlabeled:
		rtnValue = false
	}
	return rtnValue
}

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
