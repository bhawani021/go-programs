package ds

func reverseArrWithIteration(arr []int) {
	min := 0
	max := len(arr) - 1

	for min < max {
		temp := arr[min]
		arr[min] = arr[max]
		arr[max] = temp

		min++
		max--
	}
}
