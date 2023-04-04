package model

type Between struct {
	UserID    int64      `json:"user_id"`
	StartDate CustomTime `json:"start_date"`
	EndDate   CustomTime `json:"end_date"`
}
