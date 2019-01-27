package franc

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// TestNewFranc tests a constructor of Dollar object
func TestNewFranc(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	amount := rand.Int()
	d := NewFranc(amount)
	assert.Equal(t, d.GetAmount(), amount)
}

// TestMultiplication is a test to test multiplcation.
func TestMultiplication(t *testing.T) {
	five := NewFranc(5)
	assert.Equal(t, five.times(2), NewFranc(10))
	assert.Equal(t, five.times(3), NewFranc(15))
}
