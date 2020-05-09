package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Article struct {
	Title       string       `json:"article"`
	Description string       `json:"description"`
}

func getCall() {

	resp, err := http.Get("https://.....")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Convert a response body to string
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

	// Convert response body to Article struct
	var article Article
	json.Unmarshal(bodyBytes, &article)

	fmt.Printf("%+v\n", article)
}

func postCall() {
	fmt.Println("2. Performing http post")

	article := Article{"Title", "Description"}

	// Take an interface input and return slice of byte
	jsonReq, err := json.Marshal(article)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post("https://....", "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	// Covert response body to string
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

	// Covert response body to Article struct
	var articleStruct Article
	json.Unmarshal(bodyBytes, &articleStruct)

	fmt.Printf("%+v\n", articleStruct)
}

func putCall() {

	article := Article{
		Title:       "title",
		Description: "Testing description",
	}

	// Convert a interface into slice of byte
	jsonReq, err := json.Marshal(article)
	if err != nil {
		log.Fatal(jsonReq)
	}

	req, err := http.NewRequest(
		http.MethodPut,
		"https://....",
		bytes.NewBuffer(jsonReq))
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)

	// Covert response body into string
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

	// Covert response body into articleStruct struct
	var articleStruct Article
	json.Unmarshal(bodyBytes, &articleStruct)
	fmt.Printf("API Response as struct: \n%+v\n", articleStruct)
}

func deleteCall() {
	req, err := http.NewRequest(http.MethodDelete, "https://....", nil)
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	// bodyBytes
	bodyBytes, err := ioutil.ReadAll(resp.Body)

	// Covert response body into string
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

	// Convert response body to Article struct
	var articleStruct Article
	json.Unmarshal(bodyBytes, &articleStruct)

	fmt.Printf("%+v\n", articleStruct)
}

func main()  {
	getCall()
	postCall()
	putCall()
	deleteCall()
}
