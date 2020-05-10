// File read/write operations using ioutil
package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func writeFile(fileName string, data []byte) {
	err := ioutil.WriteFile(fileName, data, 0777)

	if err != nil {
		log.Fatal(err)
	}
}

func readFile(fileName string) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))
}

func main()  {
	data := []byte("This is a testing data")
	fileName := "data.txt"
	// Write file
	writeFile(fileName, data)
	// Read file
	readFile(fileName)
}