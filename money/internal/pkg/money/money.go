package money

// Money is a parent package of currencies like Dollar of Franc
type Money struct {
	amount   int
	currency string
}

// newMoney constructs a Money object
func newMoney(amount int, currency string) Money {
	return Money{amount, currency}
}

// Dollar :)
func Dollar(amount int) Money {
	return newMoney(amount, "USD")
}

// Franc :)
func Franc(amount int) Money {
	return newMoney(amount, "CHF")
}

// GetAmount returns amount of Money.
func (m Money) GetAmount() int {
	return m.amount
}

// Currency returns a string of currency
func (m Money) Currency() string {
	return m.currency
}

// Times :)
func (m Money) Times(multiplier int) Expression {
	return newMoney(m.amount*multiplier, m.currency)
}

// Plus :)
func (m Money) Plus(addend Expression) Expression {
	return NewSum(m, addend)
}

func (m Money) reduce(bank Bank, to string) Money {
	rate := bank.rate(m.currency, to)
	return newMoney(m.amount/rate, to)
}
