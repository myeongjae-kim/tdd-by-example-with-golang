package franc

// Franc is a struct indicating the curreny franc.
type Franc struct {
	amount int
}

func (d *Franc) times(multiplier int) Franc {
	return Franc{d.amount * multiplier}
}
