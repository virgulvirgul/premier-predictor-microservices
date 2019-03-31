package database

const QUERY_GET_ALL_WORKOUT_EXERCISES = "SELECT workout_date as 'date', weight, name, WorkoutExercise.exerciseID as 'exerciseId', Exercise.bodyPartID as 'bodyPartId'" +
	" FROM `Set`" +
	" INNER JOIN WorkoutExercise" +
	" ON `Set`.workoutExerciseID = WorkoutExercise.workoutExerciseID" +
	" INNER JOIN Exercise" +
	" ON WorkoutExercise.exerciseID = Exercise.exerciseID" +
	" INNER JOIN Workout" +
	" ON Workout.workoutID = WorkoutExercise.workoutID" +
	" WHERE Workout.userID = ?" +
	" AND weight IS NOT NULL" +
	" GROUP BY WorkoutExercise.workoutExerciseID" +
	" ORDER BY date DESC"

const QUERY_GET_USER_COUNT = "SELECT COUNT(*) FROM User"

const QUERY_GET_WORKOUT_BY_ID = "SELECT workoutID as 'workoutId'," +
	" workout_date as 'workoutDate'," +
	" (" +
	" SELECT json_array(GROUP_CONCAT(json_object('exercise', WorkoutExercise.exerciseID," +
	" 'exerciseName', name," +
	" 'bodyPart', bodyPartID," +
	" 'workoutExerciseId', workoutExerciseID," +
	" 'sets', (SELECT json_array(GROUP_CONCAT(json_object('setId', setID, 'reps', reps, 'weight', weight)))" +
	" FROM `Set`" +
	" WHERE `Set`.workoutExerciseID = WorkoutExercise.workoutExerciseID" +
	" )" +
	" )))" +
	" FROM WorkoutExercise" +
	" INNER JOIN Exercise" +
	" ON Exercise.exerciseID = WorkoutExercise.exerciseID" +
	" WHERE WorkoutExercise.workoutID = Workout.workoutID" +
	" ) as 'exercises'" +
	" FROM Workout" +
	" WHERE workoutID = ?;"

const QUERY_UPDATE_WORKOUT_DATE = "UPDATE Workout SET workout_date = ? WHERE workoutID = ?"

const QUERY_INSERT_MULTIPLE_SETS = "INSERT INTO `Set`(workoutExerciseID, reps, weight) VALUES "

const QUERY_INSERT_WORKOUT_EXERCISE = "INSERT INTO WorkoutExercise(exerciseID, workoutID) VALUES (?, ?)"

const QUERY_UPDATE_WORKOUT_EXERCISE = "UPDATE WorkoutExercise SET exerciseID = ? WHERE workoutExerciseID = ?"

const QUERY_UPDATE_SET = "UPDATE `Set` SET reps = ?, weight = ? WHERE setID = ?"

const QUERY_DELETE_WORKOUT_EXERCISES = "DELETE FROM WorkoutExercise WHERE workoutExerciseID IN"

const QUERY_DELETE_SETS = "DELETE FROM `Set` WHERE setID IN"

const QUERY_INSERT_WORKOUT = "INSERT INTO Workout(workout_date, userID)" +
	" VALUES (?, ?)"

const QUERY_GET_WORKOUT_LIST = "SELECT Workout.workoutID as workoutID, Workout.workout_date as workout_date, BodyPart.name as bodyPart" +
	" FROM Workout" +
	" RIGHT JOIN WorkoutExercise" +
	" On Workout.workoutID = WorkoutExercise.workoutID" +
	" RIGHT JOIN Exercise" +
	" On Exercise.exerciseID = WorkoutExercise.exerciseID" +
	" RIGHT JOIN BodyPart" +
	" On Exercise.bodyPartID = BodyPart.bodyPartID" +
	" WHERE Workout.userID = ?" +
	" AND (workout_date > DATE_SUB(NOW(), INTERVAL 120 DAY))" +
	" GROUP BY workoutID, bodyPart" +
	" ORDER BY workout_date DESC, bodyPart"

const QUERY_GET_ATTENDANCE = "SELECT workout_date as 'workoutDate', COUNT(*) as 'count'" +
	" FROM Workout" +
	" WHERE userID = ?" +
	" GROUP BY MONTH(workout_date), YEAR(workout_date)" +
	" ORDER BY workout_date DESC" +
	" LIMIT 12"
