package main

import "fmt"

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Faraaz"
	middleName = "Ahmad"
	lastName = "Permadi"
	return firstName, middleName, lastName
}

func main() {
	firstName, middleName, lastName := getCompleteName()
	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)
	fmt.Println(getCompleteName())
}