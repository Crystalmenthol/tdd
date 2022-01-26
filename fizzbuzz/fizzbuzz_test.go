package fizzbuzz

import (
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	t.Run("test FizzBuzz with 1 should get 1", func(t *testing.T) {
		expected := "1"

		actual := FizzBuzz(1)

		if actual != expected {
			t.Errorf("expected %v, but got %v ", expected, actual)
		}
	})

	t.Run("test FizzBuzz with 3 should get Fizz", func(t *testing.T) {
		expected := "Fizz"

		actual := FizzBuzz(3)

		if actual != expected {
			t.Errorf("expected %v, but got %v ", expected, actual)
		}
	})

	t.Run("test FizzBuzz with 5 should get Buzz", func(t *testing.T) {
		expected := "Buzz"

		actual := FizzBuzz(5)

		if actual != expected {
			t.Errorf("expected %v, but got %v ", expected, actual)
		}
	})

	t.Run("test FizzBuzz with 15 should get FizzBuzz", func(t *testing.T) {
		expected := "Fizz Buzz"

		actual := FizzBuzz(15)

		if actual != expected {
			t.Errorf("expected %v, but got %v ", expected, actual)
		}
	})
}
