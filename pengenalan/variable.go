package main

import "fmt"

func main() {
	var name string

	name = "Faraaz Ahmad Permadi"
	fmt.Println(name)

	name = "John Doe"
	fmt.Println(name)

	var age = 21
	fmt.Println(age) // int

	var age8 int8 = 21
	fmt.Println(age8) // int8

	var agee int16
	agee = 21
	fmt.Println(agee) // int16

	address := "Jl. Kebon Jeruk Raya No. 1"
	fmt.Println(address)

	address = "Jl. Kebon Jeruk Raya No. 2"
	fmt.Println(address)

	var (
		firstName = "Faraaz"
		lastName  = "A Permadi"
	)

	fmt.Println(firstName, lastName)
}