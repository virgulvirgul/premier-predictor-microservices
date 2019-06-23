package user

import (
	"context"
	"github.com/cshep4/premier-predictor-microservices/src/userservice/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"os"
	"testing"
)

var id = primitive.NewObjectID().Hex()

const (
	firstName = "a first name"
	surname = "a surname"
	email = "an email address"
	password = "a new password"
	score = 12345
)

func TestRepository_GetUserById(t *testing.T) {
	repository := newTestRepository(t)
	defer repository.Close()

	t.Run("returns not found if user cannot be found", func(t *testing.T) {
		cleanupDb(repository)

		user, err := repository.GetUserById(id)
		require.Error(t, err)

		assert.Equal(t, model.ErrUserNotFound, err)
		assert.Nil(t, user)
	})

	t.Run("gets the specified user", func(t *testing.T) {
		defer cleanupDb(repository)

		user := model.User{
			Id: id,
			FirstName: "Chris",
			Surname: "Shepherd",
		}

		entity, err := fromUser(user)
		require.NoError(t, err)

		createUser(entity, repository, t)

		result, err := repository.GetUserById(id)
		require.NoError(t, err)

		assert.Equal(t, &user, result)
	})
}

func TestRepository_UpdateUserInfo(t *testing.T) {
	repository := newTestRepository(t)
	defer repository.Close()

	userInfo := model.UserInfo{
		Id: id,
		FirstName: firstName,
		Surname: surname,
		Email: email,
	}

	t.Run("returns error if the user does not exist", func(t *testing.T) {
		cleanupDb(repository)

		err := repository.UpdateUserInfo(userInfo)
		require.Error(t, err)

		assert.Equal(t, model.ErrUserNotFound, err)
	})

	t.Run("updates the correct user's information", func(t *testing.T) {
		defer cleanupDb(repository)

		user := model.User{
			Id: id,
		}

		entity, err := fromUser(user)
		require.NoError(t, err)

		createUser(entity, repository, t)

		err = repository.UpdateUserInfo(userInfo)
		require.NoError(t, err)

		result, err := repository.GetUserById(id)
		require.NoError(t, err)

		assert.Equal(t, id, result.Id)
		assert.Equal(t, firstName, result.FirstName)
		assert.Equal(t, surname, result.Surname)
		assert.Equal(t, email, result.Email)
	})
}

func TestRepository_UpdatePassword(t *testing.T) {
	repository := newTestRepository(t)
	defer repository.Close()

	t.Run("returns error if the user does not exist", func(t *testing.T) {
		cleanupDb(repository)

		err := repository.UpdatePassword(id, password)
		require.Error(t, err)

		assert.Equal(t, model.ErrUserNotFound, err)
	})

	t.Run("updates the correct user's password", func(t *testing.T) {
		defer cleanupDb(repository)

		user := model.User{
			Id: id,
			FirstName: firstName,
			Password: "an old password",
		}

		entity, err := fromUser(user)
		require.NoError(t, err)

		createUser(entity, repository, t)

		err = repository.UpdatePassword(id, password)
		require.NoError(t, err)

		result, err := repository.GetUserById(id)
		require.NoError(t, err)

		assert.Equal(t, id, result.Id)
		assert.Equal(t, password, result.Password)
	})
}

func TestRepository_GetAllUsers(t *testing.T) {
	repository := newTestRepository(t)
	defer repository.Close()

	t.Run("returns empty slice if no users exist", func(t *testing.T) {
		cleanupDb(repository)

		users, err := repository.GetAllUsers()
		require.NoError(t, err)

		assert.NotNil(t, users)
		assert.Equal(t, 0, len(users))
	})

	t.Run("returns all users from db", func(t *testing.T) {
		defer cleanupDb(repository)

		id1 := primitive.NewObjectID()
		id2 := primitive.NewObjectID()
		id3 := primitive.NewObjectID()

		createUser(&userEntity{Id: &id1, Email: "1"}, repository, t)
		createUser(&userEntity{Id: &id2, Email: "2"}, repository, t)
		createUser(&userEntity{Id: &id3, Email: "3"}, repository, t)

		users, err := repository.GetAllUsers()
		require.NoError(t, err)

		assert.NotNil(t, users)
		assert.Equal(t, 3, len(users))
		assert.Equal(t, id1.Hex(), users[0].Id)
		assert.Equal(t, id2.Hex(), users[1].Id)
		assert.Equal(t, id3.Hex(), users[2].Id)
	})
}

func TestRepository_GetAllUsersByIds(t *testing.T) {
	repository := newTestRepository(t)
	defer repository.Close()

	t.Run("returns empty slice if no users exist", func(t *testing.T) {
		cleanupDb(repository)

		users, err := repository.GetAllUsers()
		require.NoError(t, err)

		assert.NotNil(t, users)
		assert.Equal(t, 0, len(users))
	})

	t.Run("returns all users with specified id", func(t *testing.T) {
		defer cleanupDb(repository)

		id1 := primitive.NewObjectID()
		id2 := primitive.NewObjectID()
		id3 := primitive.NewObjectID()

		createUser(&userEntity{Id: &id1, Email: "1"}, repository, t)
		createUser(&userEntity{Id: &id2, Email: "2"}, repository, t)
		createUser(&userEntity{Id: &id3, Email: "3"}, repository, t)

		users, err := repository.GetAllUsersByIds([]string{id1.Hex(), id3.Hex()})
		require.NoError(t, err)

		assert.NotNil(t, users)
		assert.Equal(t, 2, len(users))
		assert.Equal(t, id1.Hex(), users[0].Id)
		assert.Equal(t, id3.Hex(), users[1].Id)
	})
}

func TestRepository_IsEmailTaken(t *testing.T) {
	repository := newTestRepository(t)
	defer repository.Close()

	t.Run("returns true if a different user already has the same email", func(t *testing.T) {
		defer cleanupDb(repository)

		otherId := primitive.NewObjectID()

		createUser(&userEntity{Id: &otherId, Email: email}, repository, t)

		taken := repository.IsEmailTakenByADifferentUser(id, email)

		assert.True(t, taken)
	})

	t.Run("returns false if the specified user already has the same email", func(t *testing.T) {
		defer cleanupDb(repository)

		id := primitive.NewObjectID()

		createUser(&userEntity{Id: &id, Email: email}, repository, t)

		taken := repository.IsEmailTakenByADifferentUser(id.Hex(), email)

		assert.False(t, taken)
	})

	t.Run("returns false if the email address is not taken", func(t *testing.T) {
		cleanupDb(repository)

		taken := repository.IsEmailTakenByADifferentUser(id, email)

		assert.False(t, taken)
	})
}

func newTestRepository(t *testing.T) *repository {
	err := os.Setenv("MONGO_PORT", "27017")
	require.NoError(t, err)
	err = os.Setenv("MONGO_HOST", "localhost")
	require.NoError(t, err)
	err = os.Setenv("MONGO_SCHEME", "mongodb")
	require.NoError(t, err)

	repository, err := NewRepository()
	require.NoError(t, err)

	return repository
}

func createUser(u *userEntity, repository *repository, t *testing.T) {
	_, err := repository.
		client.
		Database(db).
		Collection(collection).
		InsertOne(
			context.Background(),
			u,
		)

	require.NoError(t, err)
}

func cleanupDb(repository *repository) {
	_, _ = repository.
		client.
		Database(db).
		Collection(collection).
		DeleteMany(
			context.Background(),
			bson.M{},
		)
}
