package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type User struct {
	Name       string
	Age        int
	City       string
}

func writeFile(fileName string, user User) {

	file, err := json.MarshalIndent(user, "", "")
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(fileName, file, 0664)
	if err != nil {
		panic(err)
	}
}

func readFile(fileName string) {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(file))
	var user User
	err = json.Unmarshal(file, &user)
	if err != nil {
		panic(err)
	}

	fmt.Println(user)

}

func main()  {
	fileName := "data.json"

	user := User{
		Name: "Bhawani",
		Age: 36,
		City: "Jaipur",
	}

	// Write operation
	writeFile(fileName, user)

	// Read operation
	readFile(fileName)
}
