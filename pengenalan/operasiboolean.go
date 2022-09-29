package main

import "fmt"

func main() {
	var nilaiAkhir = 90
	var absensi = 75

	var lulusNilaiAkhir bool = nilaiAkhir >= 80
	var lulusAbsensi bool = absensi >= 80

	var lulus bool = lulusNilaiAkhir && lulusAbsensi
	fmt.Println(lulus)

	var lulus2 bool = lulusNilaiAkhir || lulusAbsensi
	fmt.Println(lulus2)

	var lulus3 bool = !lulusNilaiAkhir
	fmt.Println(lulus3)

	var lulus4 bool = !lulusAbsensi
	fmt.Println(lulus4)
}