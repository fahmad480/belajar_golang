package main

import (
	"belajar-golang-dasar/helper"
	"fmt"
)

func main() {
	helper.SayHello("Faraaz")
	// helper.sayGoodbye("Faraaz") // error
	fmt.Println(helper.Application)
	// fmt.Println(helper.version) // error
}
