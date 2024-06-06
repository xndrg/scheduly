package timestamp

import (
	"fmt"
	"log"
	"time"
)

type Timestamp struct {
	Year  int
	Month int
	Day   int
}

func New() *Timestamp {
	const fn = "timestamp.New"

	loc, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		log.Fatal(fmt.Errorf("%s: %w", fn, err))
	}

	return &Timestamp{
		Year:  time.Now().In(loc).Year(),
		Month: int(time.Now().In(loc).Month()),
		Day:   time.Now().In(loc).Day(),
	}
}
