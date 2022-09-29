package main

import (
	"fmt"
)

func main() {
	person := map[string]string{
		"name":    "Faraaz",
		"address": "Indonesia",
	}

	person["title"] = "Programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["title"])

	var book map[string]string = make(map[string]string)
	book["title"] = "Belajar Golang"
	book["author"] = "Faraaz"
	book["ups"] = "Salah"

	fmt.Println(book)
	fmt.Println(len(book))
	
	delete(book, "ups")

	fmt.Println(book)
	fmt.Println(len(book))
}