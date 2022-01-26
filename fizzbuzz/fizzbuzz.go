package fizzbuzz

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		result := FizzBuzz(i)
		fmt.Println(result)
	}
}

func FizzBuzz(number int) string {
	if number%3 == 0 && number%5 == 0 {
		return "Fizz Buzz"
	} else if number%3 == 0 {
		return "Fizz"
	} else if number%5 == 0 {
		return "Buzz"
	}
	return fmt.Sprint(number)
}
