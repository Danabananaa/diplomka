package model

type Assets struct {
	ID          int64      `json:"id"`
	UserID      int64      `json:"user_id"`
	TypeID      int64      `json:"type_id"`
	Description string     `json:"description"`
	Amount      int64      `json:"amount"`
	Date        CustomTime `json:"date"`
}
