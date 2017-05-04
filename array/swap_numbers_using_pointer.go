package array

import "fmt"

// SwapNumbersUsingPointer ...
func SwapNumbersUsingPointer() {
	// Declare two integer type variables
	var a, b int
	fmt.Print("Please enter two numbers: ")
	fmt.Scanf("%d %d", &a, &b)

	swap(&a, &b)
	fmt.Printf("a = %d, b = %d", a, b)
}

func swap(a, b *int) {
	var temp int
	temp = *a
	*a = *b
	*b = temp
}
