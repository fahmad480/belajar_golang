package main

import "fmt"

func main() {
	name1 := "Faraaz"
	name2 := "Faraaz"

	var result bool = name1 == name2
	fmt.Println(result)

	var result2 bool = name1 != name2
	fmt.Println(result2)

	var num1 int = 10
	var num2 int = 20

	var result3 bool = num1 > num2
	fmt.Println(result3)

	var result4 bool = num1 < num2
	fmt.Println(result4)
}