package mock

import (
	"database/sql"
	. "github.com/stretchr/testify/mock"
)

type SaveWorkoutRepository struct {
	Mock
}

func (s SaveWorkoutRepository) Close() error {
	return nil
}

func (s SaveWorkoutRepository) GetTx() *sql.Tx {
	return nil
}

func (s SaveWorkoutRepository) Rollback() error {
	return nil
}

func (s SaveWorkoutRepository) Commit() error {
	return nil
}

func (s SaveWorkoutRepository) CheckErr(err error) {

}