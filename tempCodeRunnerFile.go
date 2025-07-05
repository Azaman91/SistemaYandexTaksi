// package main

// import (
// 	"fmt"
// 	"time"
// )

// func currentDayOfTheWeek() string {
// 	t := time.Now().Weekday().String()
// 	switch t {
// 	case "Monday":
// 		return "Понедельник"
// 	case "Tuesday":
// 		return "Вторник"
// 	case "Wednesday":
// 		return "Среда"
// 	case "Thursday":
// 		return "Четверг"
// 	case "Friday":
// 		return "Пятница"
// 	case "Saturday":
// 		return "Суббота"
// 	case "Sunday":
// 		return "Воскресенье"
// 	}
// 	return "ОШибка"
// }

// func dayOrNight() string {
// 	s := time.Hour.Hours() + 9.26
// 	if s >= 10 && s < 22 {
// 		return "День"
// 	}
// 	return "Ночь"
// }

// func nextFriday() int {
// 	now := time.Now()
// 	t := now.Weekday()

// 	friday := time.Friday

// 	if t < friday {
// 		return int(friday - t)
// 	} else if t > friday {
// 		return int(7 - (t - friday))
// 	}
// 	return 0
// }

// func CheckCurrentDayOfTheWeek(answer string) bool {
// 	if answer == currentDayOfTheWeek() {
// 		return true
// 	}
// 	return false
// }

// func СheckNowDayOrNight(answer string) (bool, error) {
// 	err := fmt.Errorf("исправь свой ответ, а лучше ложись поспать")

// 	if len(answer) < 4 || len(answer) > 4 {
// 		return false, err
// 	}
// 	if answer == dayOrNight() {
// 		return true, nil
// 	}

// 	return false, nil
// }

// func main() {
// 	var answer string
// 	fmt.Scan(&answer)

// 	var answer2 string
// 	fmt.Scan(&answer2)

//		fmt.Println(currentDayOfTheWeek())
//		fmt.Println(dayOrNight())
//		fmt.Println(nextFriday())
//		fmt.Println(CheckCurrentDayOfTheWeek(answer))
//		fmt.Println(СheckNowDayOrNight(answer2))
//	}
package main

import (
	"errors"
	"strings"
	"time"
	"unicode/utf8"
)

// currentDayOfTheWeek возвращает текущий день недели на русском языке
func currentDayOfTheWeek() string {
	t := TimeNow()
	switch t.Weekday() {
	case time.Monday:
		return "Понедельник"
	case time.Tuesday:
		return "Вторник"
	case time.Wednesday:
		return "Среда"
	case time.Thursday:
		return "Четверг"
	case time.Friday:
		return "Пятница"
	case time.Saturday:
		return "Суббота"
	case time.Sunday:
		return "Воскресенье"
	default:
		return ""
	}
}

// dayOrNight возвращает "День" или "Ночь" в зависимости от времени
func dayOrNight() string {
	t := TimeNow()
	hour := t.Hour()
	if hour >= 10 && hour <= 22 {
		return "День"
	}
	return "Ночь"
}

// nextFriday возвращает количество дней до следующей пятницы (включительно)
func nextFriday() int {
	t := TimeNow()
	currentWeekday := t.Weekday()

	var daysUntilFriday int
	if currentWeekday == time.Friday {
		daysUntilFriday = 0
	} else {
		daysUntilFriday = (int(time.Friday) - int(currentWeekday) + 7) % 7
	}
	return daysUntilFriday
}

// CheckCurrentDayOfTheWeek проверяет, знает ли Арсений текущий день недели
func CheckCurrentDayOfTheWeek(answer string) bool {
	correctDay := strings.ToLower(currentDayOfTheWeek())
	answerLower := strings.ToLower(strings.TrimSpace(answer))
	return correctDay == answerLower
}

// CheckNowDayOrNight проверяет, знает ли Арсений, день или ночь сейчас
// Возвращает (bool, error)
func CheckNowDayOrNight(answer string) (bool, error) {
	answerTrimmed := strings.TrimSpace(answer)
	length := utf8.RuneCountInString(answerTrimmed)

	if length != 4 && length != 5 { // ожидаем строки длиной 4 ("День") или 5 ("Ночь")
		return false, errors.New("исправь свой ответ, а лучше ложись поспать")
	}

	correctAnswer := strings.ToLower(dayOrNight())
	userAnswer := strings.ToLower(answerTrimmed)

	if correctAnswer == userAnswer {
		return true, nil
	}
	return false, nil
}
