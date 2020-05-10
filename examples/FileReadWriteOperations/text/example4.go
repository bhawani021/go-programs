package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main()  {
	fileName := "test.txt"

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}