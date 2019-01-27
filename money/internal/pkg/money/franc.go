package money

// franc is a struct indicating the curreny franc.
type franc struct {
	Money
}

func (f *franc) times(multiplier int) franc {
	return NewFranc(f.GetAmount() * multiplier)
}

// NewFranc is a constructor of Franc package.
func NewFranc(amount int) franc {
	return franc{NewMoney(amount)}
}
