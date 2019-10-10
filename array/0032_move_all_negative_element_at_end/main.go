package main

import "fmt"

// Move all negative elements to end in order with extra space allowed
func Segregate(arr []int) {
	temp := make([]int, len(arr))

	var j int;
	for i := 0; i < len(arr); i++ {
		if arr[i] >= 0 {
			temp[j] = arr[i]
			j++
		}
	}

	for i := 0; i < len(arr); i++ {
		if arr[i] < 0 {
			temp[j] = arr[i]
			j++
		}
	}

	for i := 0; i < len(arr); i++ {
		arr[i] = temp[i]
	}

}

func main()  {
	arr := []int{1, -1, 2, 3, 4, 9, -8}

	Segregate(arr)

	fmt.Println(arr)
}

