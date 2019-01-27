package franc

import "github.com/myeongjae-kim/tdd-by-example-with-golang/money/internal/pkg/money"

// franc is a struct indicating the curreny franc.
type franc struct {
	money.Money
}

func (f *franc) times(multiplier int) franc {
	return NewFranc(f.GetAmount() * multiplier)
}

// NewFranc is a constructor of Franc package.
func NewFranc(amount int) franc {
	return franc{money.NewMoney(amount)}
}
