package main

import "fmt"

func main() {
	name := "Faraaz"
	counter := 0
	
	fmt.Println(name)

	increment := func() {
		fmt.Println("Increment")
		counter++
		name = "Faraaaaaaaaaaaaaaz"
	}

	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)
}