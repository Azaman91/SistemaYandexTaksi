package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	name, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	words := strings.Split(name, " ")

	mapa := make(map[string]int)

	for _, word := range words {
		mapa[word]++
	}

	fmt.Println(mapa)

}
