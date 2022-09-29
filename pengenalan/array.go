package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Faraaz"
	names[1] = "Fauzan"
	names[2] = "Fadhil"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		90,
		80,
		70,
	}

	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	fmt.Println(len(names))
	fmt.Println(len(values))

	names[1] = "Jopan"
	fmt.Println(names[1])
}