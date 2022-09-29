package main

import "fmt"

func main() {
	type noKTP string

	var noKTPFaraaz noKTP = "1234567890"
	fmt.Println(noKTPFaraaz)
	fmt.Printf("Tipe data noKTPFaraaz %T", noKTPFaraaz)
	fmt.Println(noKTP("1234567890"))
}