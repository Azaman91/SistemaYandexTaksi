package main

import (
	"fmt"
	"sort"
)

const (
	Lexical       = "lexical"
	Length        = "length"
	FirstChar     = "firstChar"
	LastChar      = "lastChar"
	NumbersAmount = "numbersAmount"
)

// Реализуйте сортировку пузырьком
func bubbleSort(slice []string, comparator func(string, string) bool) []string {
	// Реализуйте эту функцию исходя из условий задачи
	result := make([]string, len(slice))
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice); j++ {
			if comparator(slice[i], slice[j]) {
				result = append(result, slice[i])
			}
		}
	}
	return result
}

func main() {
	// дополните эту мапу своими анонимными функциями
	comparators := map[string]func(string, string) bool{
		Lexical: func(slice, x string) bool {
			slice2 := make([]string, 0)
			if len(slice) > 0 {
				slice2 = append(slice2, slice)
				sort.Strings(slice2)
				return true
			}
			return false
		}, // Сравнение строк в лексикографическом порядке
		Length: func(slice, x string) bool {
			if len(slice) > len(x) {

			}
		}, // Сравнение строк по длине
		FirstChar: func(slice, x string) bool {
		}, // Сравнение строк по первому символу
		LastChar: func(slice, x string) bool {
		}, // Сравнение строк по последнему символу
		NumbersAmount: func(slice, x string) bool {
		}, // Сравнение строк по количеству цифр
	}

	// Считывание данных
	var n int
	fmt.Scan(&n)

	strings := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&strings[i])
	}

	var operation string
	fmt.Scan(&operation)

	if comparator, exists := comparators[operation]; exists {
		sortedStrings := bubbleSort(strings, comparator)
		for _, str := range sortedStrings {
			fmt.Print(str, " ")
		}
	}
}
