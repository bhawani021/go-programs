package main

import "fmt"

const noOfChars int = 256
var arr [noOfChars]int

func getCharCountArray(str string) {
	
	for i := 0; i < len(str); i++ {
		arr[str[i]] = arr[str[i]] + 1
	}
}

func getFirstNonRepeating(str string) int {

	getCharCountArray(str)
	
	for i := 0; i < len(str); i++ {
		if arr[str[i]] == 1 {
			return i
		}
	}

	return -1
}

func main() {
	str := "dadae"
	index := getFirstNonRepeating(str)
	
	fmt.Printf("%c", rune(str[index]))
}