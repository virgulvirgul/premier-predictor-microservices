package model

import (
	"github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	"time"
)

type User struct {
	Id              string    `json:"id,omitempty"`
	FirstName       string    `json:"firstName,omitempty"`
	Surname         string    `json:"surname,omitempty"`
	Email           string    `json:"email,omitempty"`
	Password        string    `json:"password,omitempty"`
	PredictedWinner string    `json:"predictedWinner,omitempty"`
	Score           int       `json:"score,omitempty"`
	Joined          time.Time `json:"joined,omitempty"`
	Admin           bool      `json:"admin,omitempty"`
	AdFree          bool      `json:"adFree,omitempty"`
}

func UserToGrpc(user *User) *model.User {
	return &model.User{
		Id:              user.Id,
		FirstName:       user.FirstName,
		Surname:         user.Surname,
		PredictedWinner: user.PredictedWinner,
		Score:           int32(user.Score),
	}
}

func UserFromGrpc(user *model.User) *User {
	return &User{
		Id:              user.Id,
		FirstName:       user.FirstName,
		Surname:         user.Surname,
		PredictedWinner: user.PredictedWinner,
		Score:           int(user.Score),
	}
}

