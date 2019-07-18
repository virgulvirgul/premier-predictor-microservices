package league

import (
	"context"
	"errors"
	"fmt"
	"github.com/cshep4/premier-predictor-microservices/src/leagueservice/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"os"
	"time"
)

const (
	db         = "league"
	collection = "league"
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
			name:   "users_idx",
			field:  []string{"users"},
			unique: false,
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

func (r *repository) GetLeagueByPin(pin int64) (*model.League, error) {
	var u leagueEntity

	err := r.client.
		Database(db).
		Collection(collection).
		FindOne(
			context.Background(),
			bson.D{
				{
					Key:   "_id",
					Value: pin,
				},
			},
		).
		Decode(&u)

	switch {
	case err == mongo.ErrNoDocuments:
		return nil, model.ErrLeagueNotFound
	case err != nil:
		return nil, err
	}

	return toLeague(u), nil
}

func (r *repository) GetLeaguesByUserId(id string) ([]*model.League, error) {
	users := []*model.League{}

	cur, err := r.client.
		Database(db).
		Collection(collection).
		Find(
			context.Background(),
			bson.D{
				{
					Key:   "users",
					Value: id,
				},
			},
		)

	if err != nil {
		return users, err
	}

	defer cur.Close(context.Background())

	for cur.Next(context.TODO()) {
		var l leagueEntity

		err := cur.Decode(&l)
		if err != nil {
			return nil, err
		}

		users = append(users, toLeague(l))
	}

	return users, nil
}

func (r *repository) AddLeague(league model.League) error {
	_, err := r.client.
		Database(db).
		Collection(collection).
		InsertOne(
			context.Background(),
			fromLeague(league),
		)

	return err
}

func (r *repository) JoinLeague(pin int64, id string) error {
	return r.editLeague(
		pin,
		bson.D{
			{
				Key: "$addToSet",
				Value: bson.D{
					{
						Key:   "users",
						Value: id,
					},
				},
			},
		},
	)
}

func (r *repository) LeaveLeague(pin int64, id string) error {
	return r.editLeague(
		pin,
		bson.D{
			{
				Key: "$pull",
				Value: bson.D{
					{
						Key:   "users",
						Value: id,
					},
				},
			},
		},
	)
}

func (r *repository) RenameLeague(pin int64, name string) error {
	return r.editLeague(
		pin,
		bson.D{
			{
				Key: "$set",
				Value: bson.D{
					{
						Key:   "name",
						Value: name,
					},
				},
			},
		},
	)
}

func (r *repository) editLeague(pin int64, update bson.D) error {
	res, err := r.client.
		Database(db).
		Collection(collection).
		UpdateOne(
			context.Background(),
			bson.D{
				{
					Key:   "_id",
					Value: pin,
				},
			},
			update,
		)

	if err != nil {
		return err
	}

	if res.MatchedCount == 0 {
		return model.ErrLeagueNotFound
	}

	return nil
}

func (r *repository) Ping() error {
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(5000*time.Millisecond))
	return r.client.Ping(ctx, nil)
}

func (r *repository) Close() error {
	return r.client.Disconnect(context.Background())
}
