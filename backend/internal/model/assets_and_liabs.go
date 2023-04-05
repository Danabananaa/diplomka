package model

type Assets_and_Liabs struct {
	ID          int64      `json:"id"`
	UserID      int64      `json:"user_id"`
	TypeID      int64      `json:"type_id"`
	Description string     `json:"description"`
	Amount      int64      `json:"amount"`
	Date        CustomTime `json:"date"`
}

type MergeStruct struct {
	AssetsArr      []*Assets
	LiabilitiesArr []*Liabilities
}
