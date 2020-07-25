package ds

// RearrangeArrayUsingHash arrages array in such way that array[i] = i
func RearrangeArrayUsingHash(arr []int) {
	hash := make(map[int]bool)

	// Length of array
	n := len(arr)

	for i := 0; i < n; i++ {
		if _, ok := hash[arr[i]]; !ok {
			hash[arr[i]] = true
		}
	}

	for i := 0; i < n; i++ {
		if _, ok := hash[i]; ok {
			arr[i] = i
		} else {
			arr[i] = -1
		}
	}

}

// func main() {
// 	arr := []int{-1, -1, 6, 1, 9, 3, 2, -1, 4, -1}
// 	RearrangeArrayUsingHash(arr)
// 	fmt.Println(arr) // [-1 1 2 3 4 -1 6 -1 -1 9]
// }
