package database

import "reflect"

type DbRows interface {
	Next() bool
	Scan(dest ...interface{}) error
	Close() error
}

type MockRows struct {
	hits   int
	count  int
	values [][]interface{}
}

func MockQueryResult(hits int, values [][]interface{}) MockRows {
	return MockRows{hits, -1, values}
}

func (r *MockRows) Next() bool {
	r.count++
	return r.count < r.hits
}

func (r *MockRows) Scan(args ...interface{}) error {
	for i, value := range r.values[r.count] {
		destVal := reflect.ValueOf(args[i])
		val := reflect.ValueOf(value)
		if destElem := destVal.Elem(); destElem.CanSet() {
			destElem.Set(val)
		}
	}

	return nil
}

func (r *MockRows) Close() error {
	return nil
}

type MockResult struct {
	LastInsertedId  int64
	NumRowsAffected int64
}

func (r MockResult) LastInsertId() (int64, error) {
	return r.LastInsertedId, nil
}

func (r MockResult) RowsAffected() (int64, error) {
	return r.NumRowsAffected, nil
}
