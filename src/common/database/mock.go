package database

import "database/sql"

func MockExecutionPreparedQuery(res sql.Result, err error) func([]interface{}) (sql.Result, error) {
	return func(i []interface{}) (sql.Result, error) {
		return res, err
	}
}

func MockMultipleExecutionQuery(res sql.Result, err error) func([][]interface{}) (sql.Result, error) {
	return func(i [][]interface{}) (sql.Result, error) {
		return res, err
	}
}

func MockMultipleRecordExecutionPreparedQuery(res sql.Result, err error) func(string, []interface{}) (sql.Result, error) {
	return func(str string, i []interface{}) (sql.Result, error) {
		return res, err
	}
}