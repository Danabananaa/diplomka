package model

import "time"

type Spending struct {
	ID           int64
	UserID       int64
	SpendingType int64
	Amount       int64
	Date         time.Time
}
