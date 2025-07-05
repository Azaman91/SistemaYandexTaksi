package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func tindeks1(p, m, o string) int {
	s := p + m + o
	mapa := make(map[rune]bool)

	for _, num := range s {
		mapa[num] = true
	}
	return len(mapa)
}

func tindeks2(d1, d2 int) int {
	n := d1
	n2 := d2
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	for n2 > 0 {
		sum += n2 % 10
		n2 /= 10
	}
	return sum * 64
}

func tindeks3(p string) int {
	f := rune(p[0])

	up := unicode.ToUpper(f)

	l := int(up - 'A' + 1)

	return l * 256
}

func tindeksfin(p, m, o string, d1, d2 int) {
	s := tindeks1(p, m, o) + tindeks2(d1, d2) + tindeks3(p)
	h := strconv.FormatInt(int64(s), 16)
	hs := strings.ToUpper(h)
	fmt.Println(hs[1:])
}
func main() {
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var p, m, o string
		fmt.Scan(&p, &m, &o)
		var d1, d2 int
		fmt.Scan(&d1, &d2)
		tindeksfin(p, m, o, d1, d2)
	}
}
