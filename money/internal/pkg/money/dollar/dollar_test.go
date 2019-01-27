package dollar

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// TestNewDollar tests a constructor of Dollar object
func TestNewDollar(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	amount := rand.Int()
	d := NewDollar(amount)
	assert.Equal(t, d.GetAmount(), amount)
}

// TestMultiplication tests multiplcation.
func TestMultiplication(t *testing.T) {
	five := NewDollar(5)
	assert.Equal(t, five.times(2), NewDollar(10))
	assert.Equal(t, five.times(3), NewDollar(15))
}
