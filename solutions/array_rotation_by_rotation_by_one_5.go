package solutions

func LeftRotation(arr []int, d int) {
	for i := 0; i < d; i++ {
		leftRotationByOne(arr)
	}
}

func leftRotationByOne(arr []int) {
	// Length of array
	l := len(arr)

	// Get first element of array
	temp := arr[0]
	var i int
	for i = 0; i < l-1; i++ {
		arr[i] = arr[i+1]
	}
	arr[i] = temp
}