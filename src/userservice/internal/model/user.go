package model

import (
	gen "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	_ "github.com/golang/protobuf/proto"
	"time"
)

type User struct {
	Id              string    `json:"id"`
	FirstName       string    `json:"firstName"`
	Surname         string    `json:"surname"`
	Email           string    `json:"email"`
	Password        string    `json:"password"`
	PredictedWinner string    `json:"predictedWinner"`
	Score           int       `json:"score"`
	Joined          time.Time `json:"joined"`
	Admin           bool      `json:"admin"`
	AdFree          bool      `json:"adFree"`
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
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
	Surname   string `json:"surname"`
	Email     string `json:"email"`
}

type UpdatePassword struct {
	Id              string `json:"id"`
	OldPassword     string `json:"oldPassword"`
	NewPassword     string `json:"newPassword"`
	ConfirmPassword string `json:"confirmPassword"`
}

type UserScore struct {
	Score int `json:"score"`
}
