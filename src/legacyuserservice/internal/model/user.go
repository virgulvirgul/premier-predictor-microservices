package model

import (
	gen "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	_ "github.com/golang/protobuf/proto"
)

type User struct {
	Email     string `json:"email"`
	Password  string `json:"password,omitempty"`
	FirstName string `json:"firstName"`
	Surname   string `json:"surname"`
	Joined    string `json:"joined"`
	AdFree    bool   `json:"adFree"`
	Admin     bool   `json:"admin"`
}

func LegacyUserFromGrpc(user *gen.LegacyUserResponse) *User {
	return &User{
		Email:     user.Email,
		FirstName: user.FirstName,
		Surname:   user.Surname,
		Joined:    user.Joined,
		AdFree:    user.AdFree,
		Admin:     user.Admin,
	}
}

func LegacyUserToGrpc(user *User) *gen.LegacyUserResponse {
	return &gen.LegacyUserResponse{
		Email:     user.Email,
		FirstName: user.FirstName,
		Surname:   user.Surname,
		Joined:    user.Joined,
		AdFree:    user.AdFree,
		Admin:     user.Admin,
	}
}
