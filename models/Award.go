package models

import (
	"fmt"
	"strings"
)

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

func (t *Award) GetAwardDescription() string {
	var description string
	var t_list []string

	switch t.AwardType {
	case 1:
		fallthrough
	case 2:
		for i := 0; i < len(t.Recipients); i++ {
			t_list = append(t_list, strings.TrimPrefix(t.Recipients[i].TeamKey, "frc"))
		}
		description = fmt.Sprintf("%s (%s)", t.Name, strings.Join(t_list[:], ", "))
	case 3:
		fallthrough
	case 4:
		fallthrough
	case 5:
		fallthrough
	case 6:
		fallthrough
	case 7:
		description = fmt.Sprintf("%s given to %s", t.Name, t.Recipients[0].Awardee)
	default:
		description = t.Name
	}
	return description
}