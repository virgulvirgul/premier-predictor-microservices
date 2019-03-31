package database

import (
	"database/sql"
	"log"
)

func (t *TxnDbHandler) execPreparedQuery(query string, args []interface{}) (sql.Result, error) {
	stmt, err := t.tx.Prepare(query)

	if err != nil {
		log.Fatalf("Error preparing query: %v\n", err)
		return nil, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(args...)

	if err != nil {
		log.Fatalf("Error executing query: %v\n", err)
		return nil, err
	}

	return res, err
}

func (t *TxnDbHandler) preparedWriteQuery(query string) func(args []interface{}) (sql.Result, error) {
	return func(args []interface{}) (sql.Result, error) {
		return t.execPreparedQuery(query, args)
	}
}

func (t *TxnDbHandler) multipleRecordPreparedWriteQuery(query string) func(querySuffix string, args []interface{}) (sql.Result, error) {
	return func(querySuffix string, args []interface{}) (sql.Result, error) {
		return t.execPreparedQuery(query+querySuffix, args)
	}
}

func (t *TxnDbHandler) execMultipleQueries(query string, args [][]interface{}) (sql.Result, error) {
	stmt, err := t.tx.Prepare(query)

	if err != nil {
		log.Fatalf("Error preparing query: %v\n", err)
		return nil, err
	}

	defer stmt.Close()

	var res sql.Result

	for _, a := range args {
		res, err = stmt.Exec(a...)

		if err != nil {
			return nil, err
		}
	}

	if err != nil {
		log.Fatalf("Error executing query: %v\n", err)
		return nil, err
	}

	return res, err
}

func (t *TxnDbHandler) multiplePreparedWriteQueries(query string) func(args [][]interface{}) (sql.Result, error) {
	return func(args [][]interface{}) (sql.Result, error) {
		return t.execMultipleQueries(query, args)
	}
}
