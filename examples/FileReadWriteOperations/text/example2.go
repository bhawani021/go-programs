package main

import (
	"os"
)

func writeToFile(fileName string, data string) {
	f, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if _, err := f.WriteString(data); err != nil {
		panic(err)
	}
}

func main()  {
	fileName := "test.txt"
	data := "This is testing data"

	// Write to file
	writeToFile(fileName, data)
}