package model

import (
	"time"
)

type Fixture struct {
	Id        string    `json:"id,omitempty"`
	HomeTeam  string    `json:"hTeam,omitempty"`
	AwayTeam  string    `json:"aTeam,omitempty"`
	HomeGoals int       `json:"hGoals,omitempty"`
	AwayGoals int       `json:"aGoals,omitempty"`
	Played    int       `json:"played,omitempty"`
	DateTime  time.Time `json:"dateTime,omitempty"`
	Matchday  int       `json:"matchday,omitempty"`
}
