package main

import (
	"fmt"
	"math"
	"time"
)

func (c *PriceCalculator) CalculatePrice(trip TripParameters, now time.Time, weather WeatherData, lat, lng float64) float64 {
	base := CalculateBasePrice(trip)
	timeMult := GetTimeMultiplier(now)
	weatherMult := GetWeatherMultiplier(weather)
	trafficMult := GetTrafficMultiplier(c.TrafficClient.GetTrafficLevel(lat, lng))

	finalPrice := base * timeMult * weatherMult * trafficMult

	return ApplyPriceLimits(math.Round(finalPrice))
}

func main() {
	calculator := PriceCalculator{
		TrafficClient: &RealTrafficClient{},
	}

	price := calculator.CalculatePrice(
		TripParameters{Distance: 8.5, Duration: 20},
		time.Now(),
		WeatherData{HeavyRain, 10},
		55.751244, 37.618423,
	)

	fmt.Printf("Ваша цена %.0f Руб.\n", price)
}
