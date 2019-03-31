package database

import (
	"database/sql"
)

type Db interface {
	Begin() (*sql.Tx, error)
	Close() error
	Query(query string, args ...interface{}) (*sql.Rows, error)
	Prepare(query string) (*sql.Stmt, error)
}

type DbHandler struct {
	db Db
}

func (t *DbHandler) Close() error {
	return t.db.Close()
}

func (t *DbHandler) MockDb() {
	t.db = MockDb{}
}

type MockDb struct {

}

func (m MockDb) Close() error {
	return nil
}

func (m MockDb) Begin() (*sql.Tx, error) {
	return &sql.Tx{}, nil
}

func (m MockDb) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return &sql.Rows{}, nil
}

func (m MockDb) Prepare(query string) (*sql.Stmt, error) {
	return &sql.Stmt{}, nil
}