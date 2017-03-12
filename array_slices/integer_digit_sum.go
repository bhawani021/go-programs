// Sum of integer digit
package main

import "fmt"

func main() {
	var num, sum int
	fmt.Print("Please enter a number:")
	fmt.Scanf("%d", &num)

	for num > 0 {
		sum += num % 10
		num /= 10
	}

	fmt.Println("Sum =", sum)
}
