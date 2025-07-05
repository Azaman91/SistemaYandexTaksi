package main

import (
	"fmt"
)

func main() {
	var rates = map[string]float64{
		"USD": 1.0,   // Доллар США
		"EUR": 0.92,  // Евро
		"RUB": 90.0,  // Российский рубль
		"JPY": 157.0, // Японская иена
		"CNY": 7.25,  // Китайский юань
		"GBP": 0.78,  // Британский фунт
		"KZT": 460.0, // Казахстанский тенге
		"TRY": 32.5,  // Турецкая лира
		"INR": 83.0,  // Индийская рупия
		"BRL": 5.12,  // Бразильский реал
		"AUD": 1.50,  // Австралийский доллар
		"CAD": 1.36,  // Канадский доллар
		"CHF": 0.89,  // Швейцарский франк
		"SEK": 10.8,  // Шведская крона
		"NOK": 10.5,  // Норвежская крона
	}
	for {
		fmt.Println("Добро пожаловать в конвертер валют!")
		fmt.Println("Доступные валюты для конверции:")
		fmt.Println("1. NOK")
		fmt.Println("2. RUB")
		fmt.Println("3. KZT")
		fmt.Println("4. AUD")
		fmt.Println("5. CHF")
		fmt.Println("6. SEK")
		fmt.Println("7. INR")
		fmt.Println("8. BRL")
		fmt.Println("9. GBP")
		fmt.Println("10. CAD")
		fmt.Println("11. USD")
		fmt.Println("12. JPY")
		fmt.Println("13. TRY")
		fmt.Println("14. EUR")
		fmt.Println("15. CNY")
		fmt.Print("Введите сумму в USD: ")
		var sum float64
		fmt.Scan(&sum)
		if sum <= 0 {
			fmt.Println("Сумма должна превышать 0!")
			continue
		}
		fmt.Print("Выберите номер валюты для конверции из списка выше: ")
		var num int
		fmt.Scan(&num)
		switch num {
		case 1:
			{
				fmt.Printf("%.2f USD = %.2f %s\n", float64(sum), float64(sum)*rates["NOK"], "NOK")
			}
		case 2:
			{
				fmt.Printf("%.2f USD = %.2f %s\n", float64(sum), float64(sum)*rates["RUB"], "RUB")
			}
		case 3:
			{
				fmt.Printf("%.2f USD = %.2f %s\n", float64(sum), float64(sum)*rates["KZT"], "KZT")
			}
		case 4:
			{
				fmt.Printf("%.2f USD = %.2f %s\n", float64(sum), float64(sum)*rates["AUD"], "AUD")
			}
		case 5:
			{
				fmt.Printf("%.2f USD = %.2f %s\n", float64(sum), float64(sum)*rates["CHF"], "CHF")
			}
		case 6:
			{
				fmt.Printf("%.2f USD = %.2f %s\n", float64(sum), float64(sum)*rates["SEK"], "SEK")
			}
		case 7:
			{
				fmt.Printf("%.2f USD = %.2f %s\n", float64(sum), float64(sum)*rates["INR"], "INR")
			}
		case 8:
			{
				fmt.Printf("%.2f USD = %.2f %s\n", float64(sum), float64(sum)*rates["BRL"], "BRL")
			}
		case 9:
			{
				fmt.Printf("%.2f USD = %.2f %s\n", float64(sum), float64(sum)*rates["GBP"], "GBP")
			}
		case 10:
			{
				fmt.Printf("%.2f USD = %.2f %s\n", float64(sum), float64(sum)*rates["CAD"], "CAD")
			}
		case 11:
			{
				fmt.Printf("%.2f USD = %.2f %s\n", float64(sum), float64(sum)*rates["USD"], "USD")
			}
		case 12:
			{
				fmt.Printf("%.2f USD = %.2f %s\n", float64(sum), float64(sum)*rates["JPY"], "JPY")
			}
		case 13:
			{
				fmt.Printf("%.2f USD = %.2f %s\n", float64(sum), float64(sum)*rates["TRY"], "TRY")
			}
		case 14:
			{
				fmt.Printf("%.2f USD = %.2f %s\n", float64(sum), float64(sum)*rates["EUR"], "EUR")
			}
		case 15:
			{
				fmt.Printf("%.2f USD = %.2f %s\n", float64(sum), float64(sum)*rates["CNY"], "CNY")
			}
		default:
			{
				fmt.Println("Неправильный выбор валюты!")
			}
		}
	}
}
