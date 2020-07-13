package ds

func SwapNumbersUsingPointer(first, second *int) {
	temp := *first
	*first = *second
	*second = temp
}
