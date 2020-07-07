// Find first non repeating character in a given string
// Algorithm
// Step-1: construct a character count array from the input string
// Step-2: Get the first character who's count is 1

package main

import "fmt"

// Program to print the first non-repeating character
const noOfChar = 256

func getCharCountArr(str string) [noOfChar]int {
	var arr [noOfChar]int
	// Length of string
	size := len(str)
	for i := 0; i < size; i++ {
		arr[int(str[i])] += 1
	}
	return arr
}


func GetFirstNonRepeatingChar(str string) int {
	arr := getCharCountArr(str)
	var index int
	for i := 0; i < len(str); i++ {
		if arr[str[i]] == 1 {
			index = i
			break
		}

	}
	return index
}

func main()  {
	str := "testing data"
	index := GetFirstNonRepeatingChar(str)
	fmt.Println(str[index])
	fmt.Printf("%c ", rune(str[index]))
}