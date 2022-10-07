package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Faraaz Ahmad", "Faraaz"))
	fmt.Println(strings.Contains("Faraaz Ahmad", "Budi"))

	fmt.Println(strings.Split("Faraaz Kurniawna Permadi", " "))

	fmt.Println(strings.ToLower("Faraaz Ahmad Permadi"))
	fmt.Println(strings.ToUpper("Faraaz Ahmad Permadi"))
	fmt.Println(strings.ToTitle("faraaz ahmad permadi"))

	fmt.Println(strings.Trim("      Faraaz Ahmad     ", " "))
	fmt.Println(strings.ReplaceAll("Faraaz Joko Faraaz", "Faraaz", "Budi"))
}
