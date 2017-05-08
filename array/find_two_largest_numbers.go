package array

// FindTwoLargestNumbers finds two largest numbers in given array
func FindTwoLargestNumbers(nums [5]int) (l1 int, l2 int) {

	// Check length of array.
	if len(nums) < 2 {
		return l1, l2
	}

	// Check if first largest value is less than second.
	if l1 < l2 {
		var temp int
		temp = l1
		l1 = l2
		l2 = temp
	}

	// Iterate over array using range.
	for _, v := range nums {
		if v >= l1 {
			l2 = l1
			l1 = v
		} else if v > l2 {
			l2 = v
		}
	}

	return l1, l2
}
