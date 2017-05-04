package array

import "fmt"

// IntegerDigitSum ...
func IntegerDigitSum() int {
	var num, sum int
	fmt.Print("Please enter a number:")
	fmt.Scanf("%d", &num)

	for num > 0 {
		sum += num % 10
		num /= 10
	}
	return sum
}
