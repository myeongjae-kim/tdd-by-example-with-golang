package money

// Sum :)
type Sum struct {
	augend Expression
	addend Expression
}

// NewSum constructs an object of Sum
func NewSum(augend Expression, addend Expression) Sum {
	return Sum{augend: augend, addend: addend}
}

func (s Sum) reduce(bank Bank, to string) Money {
	amount := s.augend.reduce(bank, to).amount + s.addend.reduce(bank, to).amount
	return newMoney(amount, to)
}

// Plus :)
func (s Sum) Plus(addend Expression) Expression {
	return NewSum(s, addend)
}

// Times :)
func (s Sum) Times(multiplier int) Expression {
	return NewSum(s.augend.Times(multiplier), s.addend.Times(multiplier))
}
