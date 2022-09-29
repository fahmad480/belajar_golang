package main

import "fmt"

func factorialRecursive(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorialRecursive(n-1)
}

func main() {
	fmt.Println(factorialRecursive(5))
	fmt.Println(5*4*3*2*1)
}