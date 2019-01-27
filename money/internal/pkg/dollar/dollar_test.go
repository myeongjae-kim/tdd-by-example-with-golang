package dollar

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestMultiplication is a test to test multiplcation.
func TestMultiplication(t *testing.T) {
	five := Dollar{5}
	product := five.times(2)
	assert.Equal(t, product.amount, 10)
	product = five.times(3)
	assert.Equal(t, product.amount, 15)
}

// TestEquality tests the comparison of Dollar objects
func TestEquality(t *testing.T) {
	assert.True(t, Dollar{5} == Dollar{5})
	assert.False(t, Dollar{5} == Dollar{6})
}
