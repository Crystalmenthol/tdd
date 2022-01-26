package roman

import (
	"testing"
)

func TestToRoman(t *testing.T) {

	t.Run("test ToRoman with 4", func(t *testing.T) {
		expected := "IV"

		actual := ToRoman(4)

		if actual != expected {
			t.Errorf("expected %v, but got %v ", expected, actual)
		}
	})

	t.Run("test ToRoman with 9", func(t *testing.T) {
		expected := "IX"

		actual := ToRoman(9)

		if actual != expected {
			t.Errorf("expected %v, but got %v ", expected, actual)
		}
	})

	t.Run("test ToRoman with 10", func(t *testing.T) {
		expected := "X"

		actual := ToRoman(10)

		if actual != expected {
			t.Errorf("expected %v, but got %v ", expected, actual)
		}
	})

	t.Run("test ToRoman with 19", func(t *testing.T) {
		expected := "XIX"

		actual := ToRoman(19)

		if actual != expected {
			t.Errorf("expected %v, but got %v ", expected, actual)
		}
	})

	t.Run("test ToRoman with 40", func(t *testing.T) {
		expected := "XL"

		actual := ToRoman(40)

		if actual != expected {
			t.Errorf("expected %v, but got %v ", expected, actual)
		}
	})
}
