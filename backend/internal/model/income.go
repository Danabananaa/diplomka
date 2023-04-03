package model

type Income struct {
	ID           int64
	UserID       int64  `json:"user_id"`
	IncomeTypeID int64  `json:"income_type_id"`
	Amount       int64  `json:"amount"`
	Description  string `json:"description"`
	// Формат времени, в котором мы будем хранить время, - "2006-01-02".
	// С помощью такого формата мы будем получать значение в виде строки.
	// Пример кода, как из time.Time получить строку:
	// time.Now().Format("2006-01-02").
	Date CustomTime `json:"date"`
}
