package model

type Between struct {
	UserID    int64      `json:"-"`
	StartDate CustomTime `json:"start_date"`
	EndDate   CustomTime `json:"end_date"`
}
