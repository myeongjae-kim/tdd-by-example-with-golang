package money

// Sum :)
type Sum struct {
	augend Money
	addend Money
}

// NewSum constructs an object of Sum
func NewSum(augend Money, addend Money) Sum {
	return Sum{augend: augend, addend: addend}
}

func (s Sum) reduce(bank Bank, to string) Money {
	amount := s.augend.amount + s.addend.amount
	return newMoney(amount, to)
}
