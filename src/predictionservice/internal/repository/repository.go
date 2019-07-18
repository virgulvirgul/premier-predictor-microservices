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
	db         = "prediction"
	collection = "prediction"
)

var ErrPredictionNotFound = errors.New("prediction not found")

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
			name:   "userId_idx",
			field:  []string{"userId"},
			unique: false,
		},
		{
			name:   "match_idx",
			field:  []string{"matchId"},
			unique: false,
		},
		{
			name:   "prediction_idx",
			field:  []string{"userId", "matchId"},
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

func (r *repository) GetPrediction(userId, matchId string) (*model.Prediction, error) {
	var p predictionEntity

	err := r.client.
		Database(db).
		Collection(collection).
		FindOne(
			context.Background(),
			bson.M{
				"userId":  userId,
				"matchId": matchId,
			},
		).
		Decode(&p)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, ErrPredictionNotFound
		}

		return nil, err
	}

	return toPrediction(&p), nil
}

func (r *repository) GetPredictionsByUserId(id string) ([]model.Prediction, error) {
	ctx := context.Background()

	cur, err := r.client.
		Database(db).
		Collection(collection).
		Find(
			ctx,
			bson.M{
				"userId": id,
			},
		)

	if err != nil {
		return nil, err
	}

	predictions := []model.Prediction{}

	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var m predictionEntity
		err := cur.Decode(&m)
		if err != nil {
			return nil, err
		}

		predictions = append(predictions, *toPrediction(&m))
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return predictions, nil
}

func (r *repository) UpdatePredictions(predictions []model.Prediction) error {
	opts := options.FindOneAndReplaceOptions{}
	opts.SetUpsert(true)

	for _, p := range predictions {
		err := r.client.
			Database(db).
			Collection(collection).
			FindOneAndReplace(
				context.Background(),
				bson.M{
					"userId":  p.UserId,
					"matchId": p.MatchId,
				},
				fromPrediction(&p),
				&opts,
			)
		if err.Err() != nil {
			return err.Err()
		}
	}

	return nil
}

func (r *repository) GetMatchPredictionSummary(id string) (homeWins int, draw int, awayWins int, err error) {
	ctx := context.Background()

	cur, err := r.client.
		Database(db).
		Collection(collection).
		Find(
			ctx,
			bson.M{
				"matchId": id,
			},
		)

	if err != nil {
		return
	}

	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var m predictionEntity
		err = cur.Decode(&m)
		if err != nil {
			return
		}

		if m.HomeGoals > m.AwayGoals {
			homeWins++
		} else if m.HomeGoals < m.AwayGoals {
			awayWins++
		} else {
			draw++
		}
	}

	if err = cur.Err(); err != nil {
		return
	}

	return
}

func (r *repository) Ping() error {
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(5000*time.Millisecond))
	return r.client.Ping(ctx, nil)
}

func (r *repository) Close() error {
	return r.client.Disconnect(context.Background())
}
