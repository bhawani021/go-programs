package solutions

import "fmt"

func ArrayRotationUsingTempArray(arr []int, d int) []int {
	var firstPart, secondPart, finalResult []int

	for i := 0; i < d; i++ {
		firstPart = append(firstPart, arr[i])
	}

	for j := d; j < len(arr); j++ {
		secondPart = append(secondPart, arr[j])
	}

	fmt.Println(firstPart, secondPart)

	finalResult = append(secondPart, firstPart ...)

	return finalResult
}