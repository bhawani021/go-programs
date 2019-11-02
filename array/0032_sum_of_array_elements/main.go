package main

import "fmt"

func main()  {
	arr := []int{10, 20, 30}
	sum := 0
	for _, v := range arr {
		sum += v
	}

	fmt.Println(sum)
}
