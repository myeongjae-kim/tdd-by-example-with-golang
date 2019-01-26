package money

import "testing"

// TestMultiplication is a test to test multiplcation.
func TestMultiplication(t *testing.T) {
	five := Dollar{5}
	five.times(2)
	if five.amount != 10 {
		t.Error("wrong result")
	}
}
