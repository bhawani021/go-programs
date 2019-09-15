package solutions

func GetPair(arr []int, sum int) (int, int) {
	for i := 0; i < len(arr) - 1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] + arr[j] == sum {
				return i, j
			}
		}
	}
	return 0, 0
}
