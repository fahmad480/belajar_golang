package main

import "fmt"

func main() {
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1) // [Mei Juni Juli]
	fmt.Println(len(slice1)) // 3
	fmt.Println(cap(slice1)) // 8

	slice1[0] = "Mei Lahir"
	fmt.Println(slice1) // [Mei Lahir Juli]
	fmt.Println(months) // [Januari Februari Maret April Mei Lahir Juli Agustus September Oktober November Desember]

	var slice2 = months[10:]
	fmt.Println(slice2) // [November Desember]
	fmt.Println(len(slice2)) // 2
	fmt.Println(cap(slice2)) // 2

	var slice3 = append(slice2, "Ramadhan") // [November Desember Ramadhan]
	slice3[1] = "Dzuhur" // slice2[1] juga berubah
	fmt.Println(slice3) // [November Dzuhur Ramadhan]

	fmt.Println(months) // [Januari Februari Maret April Mei Lahir Juli Agustus September Oktober Dzuhur Desember]

	var slice4 = make([]string, 2, 5)
	slice4[0] = "Faraaz"
	slice4[1] = "Fauzan"
	fmt.Println(slice4) // [Faraaz Fauzan]
	fmt.Println(len(slice4)) // 2
	fmt.Println(cap(slice4)) // 5

	slice4 = append(slice4, "Fadhil")
	fmt.Println(slice4) // [Faraaz Fauzan Fadhil]
	fmt.Println(len(slice4)) // 3
	fmt.Println(cap(slice4)) // 5

	slice5 := make([]string, len(months[0:3]), cap(months[0:3]))
	copy(slice5, months[0:3])
	fmt.Println(slice5) // [Januari Februari Maret]
	fmt.Println(len(slice5)) // 3
	fmt.Println(cap(slice5)) // 5
}