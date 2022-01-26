package fizzbuzz

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	t.Run("test ", func(t *testing.T) {
		result := FizzBuzz(33)
		assert.Equal(t, "Fizz", result)
	})
}
