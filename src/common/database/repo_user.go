package database

type UserRepository struct {
	DbHandler
	GetUserCount func() (DbRows, error)
}

func NewUserRepository() UserRepository {
	db := UserRepository{}

	db.db = Connect()

	db.GetUserCount = db.readQuery(QUERY_GET_USER_COUNT)

	return db
}
