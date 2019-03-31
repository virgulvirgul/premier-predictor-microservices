package database

import (
	"log"
)

func (t *DbHandler) doReadQuery(query string) (DbRows, error) {
	rows, err := t.db.Query(query)

	if err != nil {
		log.Fatalf("Error executing query: %v\n", err)
		return nil, err
	}

	return rows, err
}

func (t *DbHandler) readQuery(query string) func() (DbRows, error) {
	return func() (DbRows, error) {
		return t.doReadQuery(query)
	}
}

func (t *DbHandler) doPreparedReadQuery(query string, args []interface{}) (DbRows, error) {
	stmt, err := t.db.Prepare(query)

	if err != nil {
		log.Fatalf("Error preparing query: %v\n", err)
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(args...)

	if err != nil {
		log.Fatalf("Error executing query: %v\n", err)
		return nil, err
	}

	return rows, err
}

func (t *DbHandler) preparedReadQuery(query string) func(args []interface{}) (DbRows, error) {
	return func(args []interface{}) (DbRows, error) {
		return t.doPreparedReadQuery(query, args)
	}
}

func (t *DbHandler) multipleRecordPreparedReadQuery(query string) func(querySuffix string, args []interface{}) (DbRows, error) {
	return func(querySuffix string, args []interface{}) (DbRows, error) {
		return t.doPreparedReadQuery(query + querySuffix, args)
	}
}