package solutions

func ArrayCyclicallyRotate(arr []int) {
	// Store last element of array in a temp variable
	temp := arr[len(arr)-1]

	for i := len(arr)-1; i > 0; i-- {
		arr[i] = arr[i-1]
	}

	arr[0] = temp
}
