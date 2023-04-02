package model

import "time"

type Spending struct {
	ID             int64
	UserID         int64
	SpendingTypeID int64
	Amount         int64
	Description    string
	// Формат времени, в котором мы будем хранить время, - "2006-01-02".
	// С помощью такого формата мы будем получать значение в виде строки.
	// Пример кода, как из time.Time получить строку:
	// time.Now().Format("2006-01-02").
	Date time.Time
}

func fdas() {
}
