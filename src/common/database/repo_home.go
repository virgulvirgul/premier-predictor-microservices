package database

type HomeRepository struct {
	DbHandler
	GetAllWorkoutExercisesById func([]interface{}) (DbRows, error)
	GetWorkoutsById func([]interface{}) (DbRows, error)
	GetAttendanceById func([]interface{}) (DbRows, error)
}

func NewHomeRepository() HomeRepository {
	db := HomeRepository{}

	db.db = Connect()

	db.GetAllWorkoutExercisesById = db.preparedReadQuery(QUERY_GET_ALL_WORKOUT_EXERCISES)
	db.GetWorkoutsById = db.preparedReadQuery(QUERY_GET_WORKOUT_LIST)
	db.GetAttendanceById = db.preparedReadQuery(QUERY_GET_ATTENDANCE)

	return db
}
