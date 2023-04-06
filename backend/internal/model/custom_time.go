package model

import (
	"fmt"
	"time"
)

type CustomTime struct {
	time.Time
}

func (t CustomTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, t.Format("2006-01-02"))), nil
}

func (t *CustomTime) UnmarshalJSON(b []byte) error {
	parsedTime, err := time.Parse(`"2006-01-02"`, string(b))
	if err != nil {
		return err
	}
	t.Time = parsedTime
	return nil
}

func (ct *CustomTime) Scan(value interface{}) error {
	if value == nil {
		ct.Time = time.Time{}
		return nil
	}
	t, ok := value.(time.Time)
	if !ok {
		return fmt.Errorf("invalid type %T for CustomTime", value)
	}
	ct.Time = t
	return nil
}

// func (ct CustomTime) Value() (driver.Value, error) {
// 	return ct.Time.Format("2006-01-02"), nil
// }
