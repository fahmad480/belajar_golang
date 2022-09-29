package main

import "fmt"

func main() {
	const name = "Faraaz Ahmad Permadi"
	fmt.Println(name)

	const zipcode string = "12345"
	fmt.Println(zipcode)

	const no = 123
	fmt.Println(no)

	const noo int = 1234
	fmt.Println(noo)

	const tidakdiprint = "Tidak akan diprint"
	// aplikasi bisa dijalankan meskipun variabel constant tidak dipakai

	const (
		firstName = "Faraaz"
		lastName  = "A Permadi"
	)

	fmt.Println(firstName, lastName)
	
	// error
	// name = "tidak bisa diubah lagi"
	// zipcode = "tidak bisa diubah lagi"

}