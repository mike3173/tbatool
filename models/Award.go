package models

type Award struct {
	Name       string           `json:"name"`
	AwardType  int              `json:"award_type"`
	EventKey   string           `json:"event_key"`
	Recipients []AwardRecipient `json:"recipient_list"`
	Year       int              `json:"year"`
}

type AwardRecipient struct {
	TeamKey string `json:"team_key"`
	Awardee string `json:"awardee"`
}