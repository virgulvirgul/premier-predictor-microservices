package database

type LoadWorkoutRepository struct {
	DbHandler
	GetWorkoutById func([]interface{}) (DbRows, error)
}

func InjectLoadWorkoutRepository() LoadWorkoutRepository {
	db := LoadWorkoutRepository{}

	db.db = Connect()

	db.GetWorkoutById = db.preparedReadQuery(QUERY_GET_WORKOUT_BY_ID)

	return db
}