package money

// Dollar is a struct indicating the curreny dollar.
type Dollar struct {
	amount int
}

func (d *Dollar) times(multiplier int) Dollar {
	d.amount *= multiplier
	return Dollar{}
}
