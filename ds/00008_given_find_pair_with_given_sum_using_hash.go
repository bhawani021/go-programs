package ds

// FindSumOfArrayElementsWithHash ...
func FindSumOfArrayElementsWithHash(arr []int, sum int) (int, int) {
	i, j := -1, -1
	hash := make(map[int]int)
	for index := 0; index < len(arr); index++ {
		if _, ok := hash[sum-arr[i]]; ok {
			i, j = hash[sum-arr[i]], index
		} else {
			hash[arr[i]] = i
		}
	}

	return i, j
}
