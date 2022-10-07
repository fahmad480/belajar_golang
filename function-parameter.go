package main

import "fmt"

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

func main() {
	firstName := "Faraaz"
	sayHelloTo(firstName, "Ahmad")
	sayHelloTo("Budi", "Nugraha")
}
