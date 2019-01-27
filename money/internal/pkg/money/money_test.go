package money

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// TestNewMoney tests a constructor of Money object
func TestNewMoney(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	amount := rand.Int()

	money := NewMoney(amount)
	assert.Equal(t, money.GetAmount(), amount)
}

// TestEquality tests a comparison of Money objects
func TestEquality(t *testing.T) {
	assert.True(t, NewMoney(5) == NewMoney(5))
	assert.False(t, NewMoney(5) == NewMoney(6))
}
