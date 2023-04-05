package model

type Between struct {
	UserID    int64      `json:"-" db:"user_id"`
	StartDate CustomTime `json:"start_date" db:"start_date"`
	EndDate   CustomTime `json:"end_date" db:"end_date"`
}
