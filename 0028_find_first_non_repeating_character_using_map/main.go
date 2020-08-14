package main

import "fmt"

func setMapWithFrequency(str string) map[uint8]int {
	hash := make(map[uint8]int)

	for i := 0; i < len(str); i++ {
		if _, ok := hash[str[i]]; ok {
			hash[str[i]] = hash[str[i]] + 1
		} else {
			hash[str[i]] = 1
		}
	}

	return hash
}

func getFirstNonRepeatingChar(str string) int {

	hash := setMapWithFrequency(str)

	for i := 0; i < len(str); i++ {
		if hash[str[i]] == 1 {
			return i
		}
	}

	return -1
}

func main() {
	str := "dadae"
	index := getFirstNonRepeatingChar(str)

	fmt.Printf("%c", rune(str[index]))
}
