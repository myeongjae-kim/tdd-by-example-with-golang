package money

// Bank :)
type Bank struct{}

func (b Bank) reduce(source Expression, to string) Money {
	return newMoney(10, to)
}

// NewBank constructs an object of Bank
func NewBank() Bank {
	return Bank{}
}
