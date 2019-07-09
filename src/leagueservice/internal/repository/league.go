package league

import (
	"github.com/cshep4/premier-predictor-microservices/src/leagueservice/internal/model"
)

type leagueEntity struct {
	Pin   int64    `bson:"_id,omitempty"`
	Name  string   `bson:"name,omitempty"`
	Users []string `bson:"users,omitempty"`
}

func fromLeague(league model.League) *leagueEntity {
	return &leagueEntity{
		Pin:   league.Pin,
		Name:  league.Name,
		Users: league.Users,
	}
}

func toLeague(league leagueEntity) *model.League {
	return &model.League{
		Pin:   league.Pin,
		Name:  league.Name,
		Users: league.Users,
	}
}
