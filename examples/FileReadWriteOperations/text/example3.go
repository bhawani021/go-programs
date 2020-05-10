package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func readFileLines(fileName string) (map[int]string, error) {
	data := make(map[int]string)
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		return data, err
	}


	for n, line := range strings.Split(string(file), "\n") {
		data[n] = line
	}

	return data, nil

}

func main()  {
	fileName := "test.txt"
	data, err := readFileLines(fileName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(data)
}
