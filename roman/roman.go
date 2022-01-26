package roman

import (
	"fmt"
)

var num = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
var code = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

func main() {
	for i := 1; i <= 50; i++ {
		fmt.Println(ToRoman(i))
	}
}

func ToRoman(n int) string {
	res := ""
	for i := range num {
		for n >= num[i] {
			res += code[i]
			n -= num[i]
		}
	}
	return res
}
