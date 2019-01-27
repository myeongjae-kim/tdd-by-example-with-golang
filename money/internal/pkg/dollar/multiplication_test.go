package dollar

import "testing"

// TestMultiplication is a test to test multiplcation.
func TestMultiplication(t *testing.T) {
	five := Dollar{5}
	product := five.times(2)
	if product.amount != 10 {
		t.Errorf("wrong result, expected: %d, result: %d", 10, product.amount)
	}
	product = five.times(3)
	if product.amount != 15 {
		t.Errorf("wrong result, expected: %d, result: %d", 15, product.amount)
	}
}
