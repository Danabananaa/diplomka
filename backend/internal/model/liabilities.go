package model

type Liabilities struct {
	ID          int64  `json:"id"`
	UserID      int64  `json:"user_id"`
	TypeID      int64  `json:"type_id"`
	Description string `json:"description"`
	Amount      int64
	Date        CustomTime `json:"date"`
}
