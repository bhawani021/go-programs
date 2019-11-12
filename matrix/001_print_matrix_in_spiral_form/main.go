package main

import "fmt"

func main() {
	var m, n int

	// Row
	m = 4
	// Column
	n = 4

	// k - starting row index
	// m - staring column index
	var k, l int;

	matrix := [4][4]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}


	for k < m && l < n {
		/* Print the first row from the remaining rows */
		for i := l; i < n; i++ {
			fmt.Print(matrix[k][i], " ")
		}
		k++;

		/* Print the last column from the remaining columns */
		for i := k; i < m; i++ {
			fmt.Print(matrix[i][n - 1], " ")
		}
		n--;

		/* Print the last row from the remaining rows */
		if k < m {
			for i := n - 1; i >= l; i-- {
				fmt.Print(matrix[m - 1][i], " ")
			}
			m--;
		}

		/* Print the first column from the remaining columns */
		if l < n {
			for i := m - 1; i >= k; i-- {
				fmt.Print(matrix[i][l], " ");
			}
			l++;
		}
	}





}
