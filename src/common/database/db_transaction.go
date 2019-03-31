package database

import (
	"database/sql"
	"log"
)

type TxnDbRepository interface {
	GetTx() *sql.Tx
	Close() error
	Rollback() error
	Commit() error
	CheckErr(error)
}

type TxnDbHandler struct {
	DbHandler
	tx *sql.Tx
}

func (t *TxnDbHandler) init() {
	t.db = Connect()

	tx, err := t.db.Begin()
	if err != nil {
		log.Fatalf("Error starting transaction: %v\n", err)
	}

	t.tx = tx
}

func (t TxnDbHandler) GetTx() *sql.Tx {
	return t.tx
}

func (t *TxnDbHandler) Rollback() error {
	return t.tx.Rollback()
}

func (t *TxnDbHandler) Commit() error {
	return t.tx.Commit()
}

func (t *TxnDbHandler) CheckErr(err error) {
	if err != nil {
		t.Rollback()
		panic(err)
	}
}
