package money

// Money is a parent package of currencies like Dollar of Franc
type Money struct {
	amount   int
	currency string
}

// NewMoney constructs a Money object
func NewMoney(amount int, currency string) Money {
	return Money{amount, currency}
}

// GetAmount returns amount of Money.
func (m Money) GetAmount() int {
	return m.amount
}

// Currency :)
func (m Money) Currency() string {
	return m.currency
}
