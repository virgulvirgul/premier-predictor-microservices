package prediction

import (
	"context"
	"errors"
	"fmt"
	"github.com/cshep4/premier-predictor-microservices/src/common/model"
	"go.mongodb.org/mongo-driver/bson"
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
	year, month, day := time.Now().Date()
	today := time.Date(year, month, day, 0, 0, 0, 0, time.Now().Location())

	return r.getMatches(
		bson.D{
			{
				Key: "matchDate",
				Value: bson.D{
					{
						Key:   "$gte",
						Value: today,
					},
				},
			},
		},
		&options.FindOptions{
			Limit: &limit,
			Sort: bson.D{
				bson.E{Key: "matchDate", Value: 1},
			},
		},
	)
}

func (r *repository) getMatches(filter interface{}, opts *options.FindOptions) ([]*model.MatchFacts, error) {
	ctx := context.Background()

	cur, err := r.client.
		Database(db).
		Collection(collection).
		Find(
			ctx,
			filter,
			opts,
		)

	if err != nil {
		return nil, err
	}

	matches := []*model.MatchFacts{}

	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var m matchFactsEntity
		err := cur.Decode(&m)
		if err != nil {
			return nil, err
		}

		matches = append(matches, toMatchFacts(&m))
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return matches, nil
}

func (r *repository) GetMatchFacts(id string) (*model.MatchFacts, error) {
	var m matchFactsEntity

	err := r.client.
		Database(db).
		Collection(collection).
		FindOne(
			context.Background(),
			bson.M{
				"_id": id,
			},
		).
		Decode(&m)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, ErrMatchNotFound
		}

		return nil, err
	}

	return toMatchFacts(&m), nil
}

func (r *repository) Ping() error {
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(500*time.Millisecond))
	return r.client.Ping(ctx, nil)
}

func (r *repository) Close() error {
	return r.client.Disconnect(context.Background())
}
