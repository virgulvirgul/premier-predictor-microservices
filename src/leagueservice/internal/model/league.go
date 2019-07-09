package model

import (
	gen "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	_ "github.com/golang/protobuf/proto"
)

type LeagueUser struct {
	Id              string `json:"id,omitempty"`
	FirstName       string `json:"firstName,omitempty"`
	Surname         string `json:"surname,omitempty"`
	PredictedWinner string `json:"predictedWinner,omitempty"`
	Score           int    `json:"score,omitempty"`
}

func LeagueUserFromGrpc(user *gen.User) *LeagueUser {
	return &LeagueUser{
		Id:              user.Id,
		FirstName:       user.FirstName,
		Surname:         user.Surname,
		PredictedWinner: user.PredictedWinner,
		Score:           int(user.Score),
	}
}

type LeagueUserSlice []*LeagueUser

func (l LeagueUserSlice) Len() int {
	return len(l)
}

func (l LeagueUserSlice) Less(i, j int) bool {
	return l[i].Score > l[j].Score
}

func (l LeagueUserSlice) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}


type League struct {
	Pin   int64    `json:"pin,omitempty"`
	Name  string   `json:"name,omitempty"`
	Users []string `json:"users,omitempty"`
}

type LeagueOverview struct {
	LeagueName string `json:"leagueName,omitempty"`
	Pin        int64 `json:"pin,omitempty"`
	Rank       int64  `json:"rank,omitempty"`
}

type OverallLeagueOverview struct {
	Rank      int64 `json:"rank,omitempty"`
	UserCount int64 `json:"userCount,omitempty"`
}

type StandingsOverview struct {
	OverallLeagueOverview OverallLeagueOverview `json:"overallLeagueOverview,omitempty"`
	UserLeagues           []LeagueOverview      `json:"userLeagues,omitempty"`
}
