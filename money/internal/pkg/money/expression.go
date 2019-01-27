package money

// Expression of different currencies
type Expression interface {
	reduce(to string) Money
}
