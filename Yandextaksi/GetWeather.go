package main
 
type WeatherCondition int

const(
	Clear WeatherCondition = iota // Ключевое слово iota присваивает каждой константе числовое значение по порядку (0, 1, 2, 3 и т.д.)
    Rain
    HeavyRain
    Snow
)

type WeatherData struct {
    Condition WeatherCondition
    WindSpeed int
}

func GetWeatherMultiplier(weather WeatherData) float64{
	s := 1.0
	switch {
	case weather.Condition == HeavyRain:
		s += 0.2
	case weather.Condition == Snow:
		s += 0.15
	case weather.Condition == Rain:
		s += 0.125
	}
	if weather.WindSpeed > 15{
		s += 0.1
	}
	return s
}
