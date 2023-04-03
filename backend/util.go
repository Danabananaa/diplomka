package util

import (
	"fmt"
	"time"
)

type week struct {
	Monday time.Time
	Sunday time.Time
}

func FindMondayAndSunday(t time.Time) week {
	w := week{}
	weekday := t.Weekday()
	daysAgo := int(weekday) - 1 // сколько дней назад был понедельник
	if daysAgo < 0 {
		daysAgo = 6 // если воскресенье, то начинаем с прошлого понедельника
	}

	w.Monday = t.AddDate(0, 0, -daysAgo)
	w.Sunday = w.Monday.AddDate(0, 0, 6)
	return w
}

func FindFirstDay(t time.Time) time.Time {
	x := t.Month()
	fmt.Println(int(x))
	g := time.Date(int(t.Year()), x, 1, 0, 0, 0, 0, time.UTC)
	return g
}
