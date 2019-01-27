package money

// Expression of different currencies
type Expression interface {
	reduce(b Bank, to string) Money
	Plus(addend Expression) Expression
}
