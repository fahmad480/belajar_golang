package main

import "fmt"

// tipe data	| nilai minimum			| nilai maksimum		| alias
// int8			| -128 					| 127
// int16		| -32768 				| 32767
// int32		| -2147483648 			| 2147483647
// int64		| -9223372036854775808 	| 9223372036854775807
// uint8		| 0 					| 255
// uint16 		| 0 					| 65535
// uint32 		| 0 					| 4294967295
// uint64 		| 0 					| 18446744073709551615
// float32 		| 1.18e-38 				| 3.40e+38
// float64 		| 2.23e-308 			| 1.80e+308
// complex64 	| (1.40e-45+0i) 		| (3.40e+38+0i)
// complex128 	| (4.94e-324+0i) 		| (1.80e+308+0i)
// byte 		| 0 					| 255					| uint8
// rune 		| -2147483648 			| 2147483647			| int32
// uint 		| 0 					| 18446744073709551615
// int 			| -9223372036854775808 	| 9223372036854775807
// uintptr 		| 0 					| 18446744073709551615


func main() {
	var i int = 42 // int
	var f float64 = float64(i) // float64
	var u uint = uint(f) // uint

	fmt.Println(i, f, u) // 42 42 42

	fmt.Println("Satu = ", 1) // Satu = 1
	fmt.Println("Dua = ", 2) // Satu = 2
	fmt.Println("Tiga Koma Lima = ", 3.5) // Tiga Koma Lima = 3.5
}