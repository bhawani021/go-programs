package main

import (
	"fmt"
)

func setMapWithFrequency(str string) map[uint8]int {
	m := make(map[uint8]int)

	for i := 0; i < len(str); i++ {
		if _, ok := m[str[i]]; ok {
			m[str[i]] = m[str[i]] + 1
		} else {
			m[str[i]] = 1
		}
	}

	return m
}

func getFirstNonRepeatingChar(str string)  int {
	m := setMapWithFrequency(str)

	fmt.Println(m)
	index := -1
	for i := 0; i < len(str); i++ {
		if m[str[i]] == 1 {
			index = i
			break
		}
	}

	return index
}

func main()  {
	str := "data"
	index := getFirstNonRepeatingChar(str)

	if index == -1 {
		fmt.Println("Could not found non-repeating character in string")
	} else {
		fmt.Printf("%c ", rune(str[index]))
	}
}