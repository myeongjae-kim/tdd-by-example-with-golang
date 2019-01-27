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

	money := NewMoney(amount, "CUR")
	assert.Equal(t, money.GetAmount(), amount)
}

// TestEquality tests a comparison of Money objects
func TestEquality(t *testing.T) {
	assert.True(t, NewMoney(5, "CUR") == NewMoney(5, "CUR"))
	assert.False(t, NewMoney(5, "CUR") == NewMoney(6, "CUR"))
}

// TestCurrency
func TestCurrency(t *testing.T) {
	assert.True(t, "USD" == NewDollar(1).Currency())
	assert.True(t, "CHF" == NewFranc(1).Currency())
}
