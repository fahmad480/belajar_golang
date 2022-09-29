package main

import "fmt"

func main() {
	name := "Faraaz"

	switch name {
	case "Faraaz":
		fmt.Println("Hello Faraaz")
	case "Fauzan":
		fmt.Println("Hello Fauzan")
	default:
		fmt.Println("Hello Stranger")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	switch { // switch tanpa kondisi
	case name == "Faraaz":
		fmt.Println("Hello Faraaz")
	case name == "Fauzan":
		fmt.Println("Hello Fauzan")
	default:
		fmt.Println("Hello Stranger")
	}
}