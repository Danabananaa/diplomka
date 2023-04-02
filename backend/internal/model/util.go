package model

import (
	"fmt"
	"time"
)

type customTime struct {
	time.Time
}

func (t customTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, t.Format("2006-01-02"))), nil
}

func (t *customTime) UnmarshalJSON(b []byte) error {
	parsedTime, err := time.Parse(`"2006-01-02"`, string(b))
	if err != nil {
		return err
	}
	t.Time = parsedTime
	return nil
}
