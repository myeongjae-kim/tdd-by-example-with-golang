package multiplication

import "testing"

func TestMultiplication(t *testing.T) {
	five := Dollar{5}
	five.times(2)
	if five.amount != 10 {
		t.Error("wrong result")
	}
}
