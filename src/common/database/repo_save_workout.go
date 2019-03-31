package database

import "database/sql"

type SaveWorkoutRepository struct {
	TxnDbHandler
	DeleteMultipleWorkoutExercises func(string, []interface{}) (sql.Result, error)
	DeleteMultipleSets             func(string, []interface{}) (sql.Result, error)
	UpdateSets 					   func([][]interface{}) (sql.Result, error)
	UpdateWorkoutDate              func([]interface{}) (sql.Result, error)
	UpdateWorkoutExercises         func([][]interface{}) (sql.Result, error)
	InsertMultipleSets             func(string, []interface{}) (sql.Result, error)
	InsertWorkoutExercise          func([]interface{}) (sql.Result, error)
	InsertWorkout		           func([]interface{}) (sql.Result, error)
}

func InjectSaveWorkoutRepository() SaveWorkoutRepository {
	this := SaveWorkoutRepository{}

	this.init()

	this.UpdateWorkoutDate = this.preparedWriteQuery(QUERY_UPDATE_WORKOUT_DATE)

	this.InsertMultipleSets = this.multipleRecordPreparedWriteQuery(QUERY_INSERT_MULTIPLE_SETS)
	this.InsertWorkoutExercise = this.preparedWriteQuery(QUERY_INSERT_WORKOUT_EXERCISE)
	this.InsertWorkout = this.preparedWriteQuery(QUERY_INSERT_WORKOUT)

	this.UpdateWorkoutExercises = this.multiplePreparedWriteQueries(QUERY_UPDATE_WORKOUT_EXERCISE)
	this.UpdateSets = this.multiplePreparedWriteQueries(QUERY_UPDATE_SET)

	this.DeleteMultipleWorkoutExercises = this.multipleRecordPreparedWriteQuery(QUERY_DELETE_WORKOUT_EXERCISES)
	this.DeleteMultipleSets = this.multipleRecordPreparedWriteQuery(QUERY_DELETE_SETS)
	this.DeleteMultipleSets = this.multipleRecordPreparedWriteQuery(QUERY_DELETE_SETS)

	return this
}
