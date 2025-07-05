package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	name := scanner.Text()
	words := strings.Split(name, " ")

	mapa := make(map[int]bool)
	slice := make([]int, 0)

	for _, num := range words {
		num, err := strconv.Atoi(num) // преобразуем строку в числа
		if err != nil {
			continue
		}
		if !mapa[num] {
			slice = append(slice, num)
		}
		mapa[num] = true
	}

	sort.Ints(slice) // отсартируем слайс в порядке возрастания

	for i, num := range slice {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(num)
	}

}
