package main

import "time"

var pricePerMinute = GetTimeMultiplier(time.Now())

const (
	pricePerKm = 10.0
)

type TripParameters struct {
	Distance float64
	Duration float64
}

func CalculateBasePrice(tripe TripParameters) float64 {
	return tripe.Distance*pricePerKm + tripe.Duration*pricePerMinute
}
