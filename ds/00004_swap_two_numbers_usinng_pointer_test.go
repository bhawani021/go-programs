package ds

import "testing"

func TestSwapNumbersUsingPointer(t *testing.T)  {
	first, second := 10, 20

	SwapNumbersUsingPointer(&first, &second)

	if first != 20 || second != 10 {
		t.Errorf("Got first: %d, second: %d. " +
			"Expected first: %d, second: %d", first, second, 20, 20)
	}
}
