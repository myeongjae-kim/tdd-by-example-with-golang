package money

// Money is a parent package of currencies like Dollar of Franc
type Money struct {
	amount int
}

// NewMoney constructs a Money object
func NewMoney(amount int) Money {
	return Money{amount}
}

// GetAmount returns amount of Money.
func (m *Money) GetAmount() int {
	return m.amount
}
