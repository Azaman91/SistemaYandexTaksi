package main

import(
	"time"
)

func GetTimeMultiplier(t time.Time) float64{
	hours := t.Hour()
	isWeek := t.Weekday() == time.Saturday || t.Weekday() == time.Sunday
	switch{
		case hours >= 0 && hours < 5:
			return 1.5
		case hours >= 7 && hours <= 10 && !isWeek:
			return 1.3
		case isWeek:
			return 1.2
		default:
			return 1.0
		}
}
