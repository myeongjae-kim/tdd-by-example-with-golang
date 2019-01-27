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
	currency := "CUR"

	money := newMoney(amount, currency)
	assert.Equal(t, money.GetAmount(), amount)
	assert.Equal(t, money.Currency(), currency)
}

// TestEquality tests a comparison of Money objects
func TestEquality(t *testing.T) {
	assert.True(t, newMoney(5, "CUR") == newMoney(5, "CUR"))
	assert.False(t, newMoney(5, "CUR") == newMoney(6, "CUR"))
	assert.False(t, newMoney(5, "USD") == newMoney(5, "CHF"))
}

// TestCurrency
func TestCurrency(t *testing.T) {
	assert.True(t, "USD" == Dollar(1).Currency())
	assert.True(t, "CHF" == Franc(1).Currency())
}

// TestMultiplication
func TestMultiplication(t *testing.T) {
	five := newMoney(5, "CUR")
	assert.Equal(t, five.Times(2), newMoney(10, five.Currency()))
	assert.Equal(t, five.Times(3), newMoney(15, five.Currency()))
}

// TestDollar
func TestDollar(t *testing.T) {
	d := Dollar(1)

	assert.Equal(t, d.GetAmount(), 1)
	assert.Equal(t, d.Currency(), "USD")
}

// TestFranc
func TestFranc(t *testing.T) {
	f := Franc(1)

	assert.Equal(t, f.GetAmount(), 1)
	assert.Equal(t, f.Currency(), "CHF")
}
