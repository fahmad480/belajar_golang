package main

import "fmt"

func getHello(name string) (string, string) {
	if name == "" {
		return "Hello Bro", "Hello Bro"
	} else {
		return "Hello " + name, "Hello " + name
	}
}

func main() {
	result1, result2 := getHello("Faraaz")
	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(getHello("Faraaz"))
	fmt.Println(getHello(""))
}