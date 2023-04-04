package model

type Income struct {
	ID           int64  `db:"id"`
	UserID       int64  `db:"user_id" json:"user_id"`
	IncomeTypeID int64  `db:"income_type_id" json:"income_type_id"`
	Amount       int64  `db:"amount" json:"amount"`
	Description  string `db:"description" json:"description"`
	// Формат времени, в котором мы будем хранить время, - "2006-01-02".
	// С помощью такого формата мы будем получать значение в виде строки.
	// Пример кода, как из time.Time получить строку:
	// time.Now().Format("2006-01-02").
	Date CustomTime `db:"date" json:"date"`
}
