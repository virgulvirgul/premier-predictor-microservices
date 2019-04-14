package chat

import (
	"context"
	"errors"
	"fmt"
	"github.com/cshep4/premier-predictor-microservices/src/chatservice/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"os"
	"time"
)

const (
	db               = "chat"
	collection       = "chat"
	collectionPrefix = "chat-"
)

var ErrChatNotFound = errors.New("chat not found")
var ErrMessageAlreadyExists = errors.New("message already exists")

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
	mongoUri = fmt.Sprintf("%s/?retryWrites=true&connectTimeoutMS=500", mongoUri)

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
				{Key: "users.id", Value: bsonx.Int64(1)},
			},
			Options: options.Index().
				SetName("users_idx").
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

func (r *repository) getChatCollection(chatId string) string {
	return collectionPrefix + chatId
}

func (r *repository) GetChatById(chatId string) (*model.Chat, error) {
	var c chatEntity

	err := r.client.
		Database(db).
		Collection(collection).
		FindOne(
			context.Background(),
			bson.M{
				"_id": chatId,
			},
		).
		Decode(&c)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, ErrChatNotFound
		}

		return nil, err
	}

	return toChat(c), nil
}

func (r *repository) CreateChat(chatId, userId string) error {
	u := []chatUser{
		{
			Id: userId,
		},
	}

	_, err := r.client.
		Database(db).
		Collection(collection).
		InsertOne(
			context.Background(),
			bson.M{
				"_id":   chatId,
				"users": u,
			},
		)

	return err
}

func (r *repository) JoinChat(chatId, userId string) error {
	_, err := r.client.
		Database(db).
		Collection(collection).
		UpdateOne(
			context.Background(),
			bson.M{
				"_id": chatId,
			},
			bson.M{
				"$push": bson.M{
					"users": chatUser{
						Id: userId,
					},
				},
			},
		)

	return err
}

func (r *repository) LeaveChat(chatId, userId string) error {
	_, err := r.client.
		Database(db).
		Collection(collection).
		UpdateOne(
			context.Background(),
			bson.M{
				"_id": chatId,
			},
			bson.M{
				"$pull": bson.M{
					"users": bson.M{
						"id": userId,
					},
				},
			},
		)

	return err
}

func (r *repository) GetLatestMessages(chatId string) ([]model.Message, error) {
	return r.getMessages(chatId, bson.M{})
}

func (r *repository) GetPreviousMessages(chatId, messageId string) ([]model.Message, error) {
	id, err := primitive.ObjectIDFromHex(messageId)
	if err != nil {
		return nil, err
	}

	return r.getMessages(chatId, bson.M{
		"_id" : bson.M{
			"$lt": id,
		},
	})
}

func (r *repository) getMessages(chatId string, filter interface{}) ([]model.Message, error) {
	ctx := context.Background()

	limit := int64(50)
	opts := &options.FindOptions{
		Limit: &limit,
		Sort: bson.D{
			bson.E{"_id", -1},
		},
	}
	cur, err := r.client.
		Database(db).
		Collection(r.getChatCollection(chatId)).
		Find(
			ctx,
			filter,
			opts,
		)

	if err != nil {
		return nil, err
	}

	var messages []model.Message

	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var m message
		err := cur.Decode(&m)
		if err != nil {
			return nil, err
		}

		messages = append(messages, *toMessage(m))
	}

	for i := len(messages)/2-1; i >= 0; i-- {
		opp := len(messages)-1-i
		messages[i], messages[opp] = messages[opp], messages[i]
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return messages, nil
}

func (r *repository) SaveReadReceipt(readReceipt model.ReadReceipt) error {
	id, err := primitive.ObjectIDFromHex(readReceipt.MessageId)
	if err != nil {
		return err
	}

	_, err = r.client.
		Database(db).
		Collection(collection).
		UpdateOne(
			context.Background(),
			bson.M{
				"_id":      readReceipt.ChatId,
				"users.id": readReceipt.SenderId,
			},
			bson.M{
				"$set": bson.M{
					"users.$.lastReadMessage": id,
					"users.$.readTime": readReceipt.DateTime.Round(time.Second),
				},
			},
		)

	return err
}

func (r *repository) SaveMessage(message model.Message) (string, error) {
	if message.Id != "" {
		return "", ErrMessageAlreadyExists
	}

	msg, err := fromMessage(message)
	if err != nil {
		return "", err
	}

	_, err = r.client.
		Database(db).
		Collection(r.getChatCollection(message.ChatId)).
		InsertOne(
			context.Background(),
			msg,
		)

	if err != nil {
		return "", err
	}

	return msg.Id.Hex(), nil
}

func (r *repository) Ping() error {
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(500*time.Millisecond))
	return r.client.Ping(ctx, &readpref.ReadPref{})
}

func (r *repository) Close() error {
	return r.client.Disconnect(context.Background())
}
