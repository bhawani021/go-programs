package ds

func reverseArrWithRecursion(arr []int, min, max int) {
	if min >= max {
		return
	}

	temp := arr[min]
	arr[min] = arr[max]
	arr[max] = temp

	reverseArrWithRecursion(arr, min+1, max-1)
}
