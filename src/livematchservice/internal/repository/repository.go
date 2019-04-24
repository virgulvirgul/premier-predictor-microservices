package live

import (
	"context"
	"errors"
	"fmt"
	"github.com/cshep4/premier-predictor-microservices/src/common/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"os"
	"time"
)

const (
	db         = "liveMatch"
	collection = "liveMatch"
)

var ErrMatchNotFound = errors.New("match not found")

var limit = int64(20)

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
	_, err := r.client.
		Database(db).
		Collection(collection).
		Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys: bsonx.Doc{
				{Key: "matchDate", Value: bsonx.Int64(1)},
			},
			Options: options.Index().
				SetName("matchDate_idx").
				SetUnique(false).
				SetSparse(true).
				SetBackground(true),
		},
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetUpcomingMatches() ([]*model.MatchFacts, error) {
	panic("implement me")
}

func (r *repository) GetMatchFacts(id string) (*model.MatchFacts, error) {
	panic("implement me")
}

func (r *repository) Ping() error {
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(500*time.Millisecond))
	return r.client.Ping(ctx, nil)
}

func (r *repository) Close() error {
	return r.client.Disconnect(context.Background())
}
