package model

import (
	gen "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	_ "github.com/golang/protobuf/proto"
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

func UserToGrpc(user *User) *gen.User {
	return &gen.User{
		Id:              user.Id,
		FirstName:       user.FirstName,
		Surname:         user.Surname,
		PredictedWinner: user.PredictedWinner,
		Score:           int32(user.Score),
	}
}

func UserFromGrpc(user *gen.User) *User {
	return &User{
		Id:              user.Id,
		FirstName:       user.FirstName,
		Surname:         user.Surname,
		PredictedWinner: user.PredictedWinner,
		Score:           int(user.Score),
	}
}

type UserInfo struct {
	Id        string `json:"id,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	Surname   string `json:"surname,omitempty"`
	Email     string `json:"email,omitempty"`
}

type UpdatePassword struct {
	Id              string `json:"id,omitempty"`
	OldPassword     string `json:"oldPassword,omitempty"`
	NewPassword     string `json:"newPassword,omitempty"`
	ConfirmPassword string `json:"confirmPassword,omitempty"`
}

type UserScore struct {
	Score int `json:"score,omitempty"`
}
