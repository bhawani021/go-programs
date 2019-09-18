package solutions


func ZeroSumSubArray(arr []int) bool {
	var sum int
	for _, v := range arr {
		sum += v
		if sum == 0 {
			return true
		}
	}

	return false
}

