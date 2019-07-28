package model

import (
	"github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	"github.com/golang/protobuf/ptypes"
	"time"
)

type Fixture struct {
	Id        string    `json:"id"`
	HomeTeam  string    `json:"hTeam"`
	AwayTeam  string    `json:"aTeam"`
	HomeGoals *int      `json:"hGoals,omitempty"`
	AwayGoals *int      `json:"aGoals,omitempty"`
	Played    int       `json:"played"`
	DateTime  time.Time `json:"dateTime"`
	Matchday  int       `json:"matchday"`
}

func FixtureFromGrpc(match *model.Match) Fixture {
	t, _ := ptypes.Timestamp(match.DateTime)

	f := Fixture{
		Id:        match.Id,
		HomeTeam:  match.HTeam,
		AwayTeam:  match.ATeam,
		Played:    int(match.Played),
		DateTime:  t,
		Matchday:  int(match.Matchday),
	}

	if f.Played != 0 {
		hGoals := int(match.HGoals)
		aGoals := int(match.AGoals)
		f.HomeGoals = &hGoals
		f.AwayGoals = &aGoals
	}

	return f
}
