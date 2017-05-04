package array

import "fmt"

// AddTwoArrays ...
func AddTwoArrays() {
	a := [5]int{10, 20, 30, 40, 50}
	b := [5]int{100, 200, 300, 400, 500}

	var c [5]int

	for i := 0; i < 5; i++ {
		c[i] = a[i] + b[i]
	}

	fmt.Println(c)
}
