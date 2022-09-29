package main

import "fmt"

func main() {
	name := "Faraaz"

	if name == "Faraaz" {
		fmt.Println("Hello Faraaz")
	} else if name == "Fauzan" {
		fmt.Println("Hello Fauzan")
	} else {
		fmt.Println("Hello Stranger")
	}

	length := len(name)

	if length > 5 {
		fmt.Println("Terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}