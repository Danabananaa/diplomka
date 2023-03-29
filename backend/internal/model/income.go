package model

import "time"

type Income struct {
	ID         int64
	UserID     int64
	IncomeType int64
	Amount     int64
	Date       time.Time
}
