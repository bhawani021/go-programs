package array

// SumWithPointer ...
func SumWithPointer(nums *[5]int) int {
	var sum int
	for _, v := range nums {
		sum += v
	}
	return sum
}
