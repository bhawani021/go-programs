package main

import "fmt"

func maxRepeatingChar(str string) string {
	// Set first char with max count
	maxRepeatedChar := fmt.Sprintf("%c", str[0])

	var maxCount int

	for i := 0; i < len(str); i++ {
		count := 1

		for j := i+1; j < len(str); j++ {
			if str[i] != str[j] {
				break
			}
			count++
		}

		if count > maxCount {
			maxCount = count
			maxRepeatedChar = fmt.Sprintf("%c", str[i])
		}
	}

	return maxRepeatedChar
}

func maxRepeating(str string) string {
	// Set first char as max count
	maxRepeatedChar := fmt.Sprintf("%c", str[0])

	maxCount := 0
	count := 1

	for i := 0; i < len(str); i++ {
		// if Char matches with next
		if i < len(str)-1 && str[i] == str[i+1] {
			count++
		} else {
			if count > maxCount {
				maxCount = count
				maxRepeatedChar = fmt.Sprintf("%c", str[i])
			}
			count = 1
		}
	}

	return maxRepeatedChar
}

func main() {
	str := "aaaabbaaccdeeeeeeee"
	max := maxRepeating(str)
	fmt.Println(max)
}