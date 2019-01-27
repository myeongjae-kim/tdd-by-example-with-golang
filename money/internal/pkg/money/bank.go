package money

// Bank :)
type Bank struct {
	rates map[pair]int
}

func (b Bank) reduce(source Expression, to string) Money {
	return source.reduce(b, to)
}

// NewBank constructs an object of Bank
func NewBank() Bank {
	return Bank{make(map[pair]int)}
}

func (b Bank) addRate(from string, to string, rate int) {
	b.rates[pair{from, to}] = rate
}

func (b Bank) rate(from string, to string) int {
	if from == to {
		return 1
	}

	return b.rates[pair{from, to}]
}
