package money

// Dollar is a struct indicating the curreny dollar.
type dollar struct {
	Money
}

func (d *dollar) times(multiplier int) dollar {
	return NewDollar(d.GetAmount() * multiplier)
}

// NewDollar is a constructor of Dollar package.
func NewDollar(amount int) dollar {
	return dollar{NewMoney(amount)}
}
