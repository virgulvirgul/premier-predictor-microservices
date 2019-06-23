package user

import (
	"errors"
	"github.com/cshep4/premier-predictor-microservices/src/userservice/internal/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type userEntity struct {
	Id              *primitive.ObjectID `bson:"_id,omitempty"`
	FirstName       string              `bson:"firstName,omitempty"`
	Surname         string              `bson:"surname,omitempty"`
	Email           string              `bson:"email,omitempty"`
	Password        string              `bson:"password,omitempty"`
	PredictedWinner string              `bson:"predictedWinner,omitempty"`
	Score           int                 `bson:"score,omitempty"`
	Joined          time.Time           `bson:"joined,omitempty"`
	Admin           bool                `bson:"admin,omitempty"`
	AdFree          bool                `bson:"adFree,omitempty"`
}

func fromUser(user model.User) (*userEntity, error) {
	id, err := primitive.ObjectIDFromHex(user.Id)
	if err != nil {
		return nil, errors.New("cannot create objectId")
	}

	return &userEntity{
		Id: &id,
		FirstName: user.FirstName,
		Surname: user.Surname,
		Email: user.Email,
		Password: user.Password,
		PredictedWinner: user.PredictedWinner,
		Score: user.Score,
		Joined: user.Joined,
		Admin: user.Admin,
		AdFree: user.AdFree,
	}, nil
}

func toUser(user userEntity) *model.User {
	return &model.User{
		Id: user.Id.Hex(),
		FirstName: user.FirstName,
		Surname: user.Surname,
		Email: user.Email,
		Password: user.Password,
		PredictedWinner: user.PredictedWinner,
		Score: user.Score,
		Joined: user.Joined,
		Admin: user.Admin,
		AdFree: user.AdFree,
	}
}
