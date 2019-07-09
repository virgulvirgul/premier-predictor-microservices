//go:generate mockgen -destination=./mocks/mock_time.go -package=timermocks github.com/cshep4/premier-predictor-microservices/src/common/timer Time

package timer

import "time"

type Time interface {
	Now() time.Time
}

type t struct {

}

func NewTime() Time {
	return t{}
}

func (t) Now() time.Time {
	return time.Now()
}
