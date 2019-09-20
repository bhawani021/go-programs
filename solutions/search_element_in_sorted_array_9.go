package solutions

func BinarySearch(arr []int, needle int) int {
	low := 0
	high := len(arr)

	for {
		if low < high {
			mid := (low + high)/2

			if arr[mid] == needle {
				return mid
			}

			if arr[mid] < needle {
				low = mid+1
			} else {
				high = mid
			}
		} else {
			break
		}
	}
	return -1
}
