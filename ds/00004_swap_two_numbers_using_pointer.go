package ds

// SwapNumbersUsingPointer ...
func SwapNumbersUsingPointer(first, second *int) {
	temp := *first
	*first = *second
	*second = temp
}
