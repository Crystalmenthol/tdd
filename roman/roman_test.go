package roman

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToRoman(t *testing.T) {

	t.Run("expect 9", func(t *testing.T) {
		result := ToRoman(9)

		assert.Equal(t, "IX", result)
	})
}
