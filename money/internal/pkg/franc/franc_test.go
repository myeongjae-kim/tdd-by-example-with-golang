package franc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestMultiplication is a test to test multiplcation.
func TestMultiplication(t *testing.T) {
	five := Franc{5}
	assert.Equal(t, five.times(2), Franc{10})
	assert.Equal(t, five.times(3), Franc{15})
}
