package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func writeCSV(fileName string, data [][]string) {
	// Create a named file with mode 0666.
	// truncating if it already exists
	// on successful, methods on returned *File can be used for I/O
	file, err := os.Create("student.csv")
	if err != nil {
		log.Fatalf("failed to create file: %s", err)
	}

	// Write data to csv file
	csvWriter := csv.NewWriter(file)
	_ = csvWriter.WriteAll(data)
	// OR
	//for _, v := range data {
	//	_ = csvWriter.Write(v)
	//}
	csvWriter.Flush()
	_ = file.Close()
}

func readCSV(fileName string) {
	// Open opens the named file for reading
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	// Read data from csv
	r := csv.NewReader(file)

	// Iterate through for loop
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(record)
	}
	// or
	//fmt.Println(r.ReadAll())
	_ = file.Close()
}

func main() {
	fileName := "./student.csv"

	// Write CSV file
	data := [][]string{
		{"Name", "Age", "City"},
		{"John", "33", "Jaipur"},
		{"Raj", "21", "Delhi"},
	}
	fmt.Println("CSV write operation started...")
	writeCSV(fileName, data)
	fmt.Println("CSV write operation ended...")

	// Read CSV file
	fmt.Println("Start Reading CSV file...")
	readCSV(fileName)
	fmt.Println("Read operation completed...")
}