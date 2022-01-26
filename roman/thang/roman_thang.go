package main

import (
	"fmt"
	"strings"
)

func ten(x int) string {
	return ""
}

func roman(x int) string {
	r := []string{}
	if x >= 9 {
		if x == 9 {
			r = append(r, "IX")
		} else {
			r = append(r, "X")
		}
		x -= 10
	}
	if x >= 4 {
		if x == 4 {
			r = append(r, "IV")
		} else {
			r = append(r, "V")
		}
		x -= 5
	}
	for i := 0; i < x; i++ {
		r = append(r, "I")
	}
	return strings.Join(r, "")
}

func main() {
	for i := 0; i < 21; i++ {
		fmt.Println(roman(i))
	}
}
