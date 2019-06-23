package user

import (
	"context"
	"errors"
	"fmt"
	"github.com/cshep4/premier-predictor-microservices/src/userservice/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"os"
	"time"
)

const (
	db         = "user"
	collection = "user"
)

var (
	ErrCannotCreateObjectId = errors.New("cannot create objectId")
)

type repository struct {
	client *mongo.Client
}

func NewRepository() (*repository, error) {
	username := os.Getenv("MONGO_USERNAME")
	password := os.Getenv("MONGO_PASSWORD")
	port := os.Getenv("MONGO_PORT")

	mongoUri := fmt.Sprintf("%s://", os.Getenv("MONGO_SCHEME"))
	if username != "" && password != "" {
		mongoUri = fmt.Sprintf("%s%s:%s@", mongoUri, username, password)
	}
	mongoUri = mongoUri + os.Getenv("MONGO_HOST")
	if port != "" {
		mongoUri = fmt.Sprintf("%s:%s", mongoUri, port)
	}

	c, err := mongo.NewClient(options.Client().ApplyURI(mongoUri))
	if err != nil {
		return nil, err
	}

	if err := c.Connect(context.Background()); err != nil {
		return nil, err
	}

	r := repository{
		client: c,
	}

	if err := r.Ping(); err != nil {
		return nil, err
	}

	if err := r.ensureIndexes(); err != nil {
		return nil, err
	}

	return &r, nil
}

func (r *repository) ensureIndexes() error {
	idxs := []struct {
		name   string
		field  []string
		unique bool
	}{
		{
			name:   "email_idx",
			field:  []string{"email"},
			unique: true,
		},
	}

	for _, i := range idxs {
		var doc bsonx.Doc
		for _, f := range i.field {
			doc = append(doc, bsonx.Elem{Key: f, Value: bsonx.Int64(1)})
		}

		opts := options.Index().
			SetName(i.name).
			SetUnique(i.unique).
			SetSparse(false).
			SetBackground(true)

		_, err := r.client.
			Database(db).
			Collection(collection).
			Indexes().CreateOne(
			context.Background(),
			mongo.IndexModel{
				Keys:    doc,
				Options: opts,
			},
		)

		if err != nil {
			return err
		}
	}

	return nil
}

func (r *repository) GetUserById(id string) (*model.User, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, ErrCannotCreateObjectId
	}

	var u userEntity

	err = r.client.
		Database(db).
		Collection(collection).
		FindOne(
			context.Background(),
			bson.D{
				{
					Key:   "_id",
					Value: objectId,
				},
			},
		).
		Decode(&u)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, model.ErrUserNotFound
		}

		return nil, err
	}

	return toUser(u), nil
}

func (r *repository) UpdateUserInfo(userInfo model.UserInfo) error {
	objectId, err := primitive.ObjectIDFromHex(userInfo.Id)
	if err != nil {
		return ErrCannotCreateObjectId
	}

	result, err := r.client.
		Database(db).
		Collection(collection).
		UpdateOne(
			context.Background(),
			bson.D{
				{
					Key:   "_id",
					Value: objectId,
				},
			},
			bson.D{
				{
					Key: "$set",
					Value: bson.D{
						{
							Key:   "firstName",
							Value: userInfo.FirstName,
						},
						{
							Key:   "surname",
							Value: userInfo.Surname,
						},
						{
							Key:   "email",
							Value: userInfo.Email,
						},
					},
				},
			},
		)

	if err != nil {
		return err
	}

	if result.ModifiedCount == 0 {
		return model.ErrUserNotFound
	}

	return nil
}

func (r *repository) UpdatePassword(id, password string) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ErrCannotCreateObjectId
	}

	result, err := r.client.
		Database(db).
		Collection(collection).
		UpdateOne(
			context.Background(),
			bson.D{
				{
					Key:   "_id",
					Value: objectId,
				},
			},
			bson.D{
				{
					Key: "$set",
					Value: bson.D{
						{
							Key:   "password",
							Value: password,
						},
					},
				},
			},
		)

	if err != nil {
		return err
	}

	if result.ModifiedCount == 0 {
		return model.ErrUserNotFound
	}

	return nil
}

func (r *repository) GetAllUsers() ([]*model.User, error) {
	return r.findUsers(bson.D{})
}

func (r *repository) GetAllUsersByIds(ids []string) ([]*model.User, error) {
	var objectIds []primitive.ObjectID

	for _, id := range ids {
		objectId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return []*model.User{}, ErrCannotCreateObjectId
		}

		objectIds = append(objectIds, objectId)
	}

	return r.findUsers(
		bson.D{
			{
				Key: "_id",
				Value: bson.D{
					{
						Key:   "$in",
						Value: objectIds,
					},
				},
			},
		},
	)
}

func (r *repository) findUsers(filter bson.D) ([]*model.User, error) {
	users := []*model.User{}

	cur, err := r.client.
		Database(db).
		Collection(collection).
		Find(
			context.Background(),
			filter,
		)

	if err != nil {
		return users, err
	}

	defer cur.Close(context.Background())

	for cur.Next(context.TODO()) {
		var u userEntity

		err := cur.Decode(&u)
		if err != nil {
			return nil, err
		}

		users = append(users, toUser(u))
	}

	return users, nil
}

func (r *repository) IsEmailTakenByADifferentUser(id, email string) bool {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return true
	}

	var u userEntity

	err = r.client.
		Database(db).
		Collection(collection).
		FindOne(
			context.Background(),
			bson.D{
				{
					Key:   "email",
					Value: email,
				},
				{
					Key:   "_id",
					Value: bson.D{
						{
							Key:   "$ne",
							Value: objectId,
						},
					},
				},
			},
		).
		Decode(&u)

	if err != nil && err == mongo.ErrNoDocuments{
		return false
	}

	return true
}

func (r *repository) Ping() error {
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(500*time.Millisecond))
	return r.client.Ping(ctx, nil)
}

func (r *repository) Close() error {
	return r.client.Disconnect(context.Background())
}
