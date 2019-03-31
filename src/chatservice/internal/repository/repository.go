package chat

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"os"
	"time"
)

const (
	db         = "chat"
	collection = "chat"
	collectionPrefix = "chat-"
)

type Repository interface {
	SaveMessage() error
}

type repository struct {
	client *mongo.Client
}

func NewRepository() (Repository, error) {
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
	mongoUri = fmt.Sprintf("%s/%s?retryWrites=true&connectTimeoutMS=500", mongoUri, os.Getenv("MONGO_DATABASE"))

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
		field  string
		unique bool
	}{
		{
			name:   "users_idx",
			field:  "users.id",
			unique: false,
		},
	}

	indexes := r.client.
		Database(db).
		Collection(collection).
		Indexes()

	for _, i := range idxs {
		_, err := indexes.CreateOne(
			context.Background(),
			mongo.IndexModel{
				Keys: bsonx.Doc{
					{Key: i.field, Value: bsonx.Int64(1)},
				},
				Options: options.Index().
					SetName(i.name).
					SetUnique(i.unique).
					SetSparse(true).
					SetBackground(true),
			},
		)

		if err != nil {
			return err
		}
	}

	return nil
}

func (r *repository) getLeagueCollection(chatId string) string {
	return collectionPrefix + chatId
}

func (r *repository) SaveMessage() error {
	panic("implement me")
}

func (r *repository) Ping() error {
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(500*time.Millisecond))
	return r.client.Ping(ctx, &readpref.ReadPref{})
}

func (r *repository) Close() error {
	return r.client.Disconnect(context.Background())
}
