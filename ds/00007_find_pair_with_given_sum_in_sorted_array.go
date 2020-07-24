package ds

// FindPairWithGivenSumInSortedArray ...
func FindPairWithGivenSumInSortedArray(arr []int, sum int) (int, int) {
	i, j := -1, -1

	min, max := 0, len(arr)-1

	for min < max {
		if arr[min]+arr[max] == sum {
			i, j = min, max
			break
		}

		if arr[min]+arr[max] > sum {
			max--
		} else {
			min++
		}
	}

	return i, j
}
