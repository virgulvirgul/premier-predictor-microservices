package model

import (
	_ "github.com/golang/protobuf/proto"
	"time"
)

type FixturePrediction struct {
	Id           string    `json:"id,omitempty"`
	PredictionId string    `json:"predictionId,omitempty"`
	HomeTeam     string    `json:"hTeam,omitempty"`
	AwayTeam     string    `json:"aTeam,omitempty"`
	HomeGoals    int       `json:"hGoals,omitempty"`
	AwayGoals    int       `json:"aGoals,omitempty"`
	HomeResult   int       `json:"hResult,omitempty"`
	AwayResult   int       `json:"aResult,omitempty"`
	Played       int       `json:"played,omitempty"`
	DateTime     time.Time `json:"dateTime,omitempty"`
	Matchday     int       `json:"matchday,omitempty"`
}

type TeamForm struct {
	Forms []*TeamMatchResult `protobuf:"bytes,1,rep,name=forms,proto3" json:"forms,omitempty"`
}

type TeamMatchResult struct {
	Result   Result   `json:"result,omitempty"`
	Score    string   `json:"score,omitempty"`
	Opponent string   `json:"opponent,omitempty"`
	Location Location `json:"location,omitempty"`
}

type Result string

const (
	WIN  Result = "WIN"
	DRAW Result = "DRAW"
	LOSS Result = "LOSS"
)

type Location string

const (
	HOME Location = "HOME"
	AWAY Location = "AWAY"
)

type PredictorData struct {
	Predictions []FixturePrediction  `json:"predictions,omitempty"`
	Forms       map[string]*TeamForm `json:"forms,omitempty"`
}
