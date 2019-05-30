package model

import (
	"github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	"github.com/golang/protobuf/ptypes"
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

func FixtureFromGrpc(match *model.Match) Fixture {
	t, _ := ptypes.Timestamp(match.DateTime)

	return Fixture{
		Id:        match.Id,
		HomeTeam:  match.HTeam,
		AwayTeam:  match.ATeam,
		HomeGoals: int(match.HGoals),
		AwayGoals: int(match.AGoals),
		Played:    int(match.Played),
		DateTime:  t,
		Matchday:  int(match.Matchday),
	}
}
