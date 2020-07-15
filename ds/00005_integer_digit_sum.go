package ds

func GetIntegerDigitSum(num int) int {
	var sum int
	for num > 0 {
		sum += num%10
		num /= 10
	}

	return sum
}
