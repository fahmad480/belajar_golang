package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func sumAllwithString(word, separator string, numbers ...int) string {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return word + separator + fmt.Sprint(total)
}

func main() {
	total := sumAll(10, 10, 10, 10)
	fmt.Println(total)
	fmt.Println(sumAll(10, 10, 10, 10, 10, 10, 10, 10, 10, 10))
	fmt.Println(sumAll())
	
	fmt.Println(sumAllwithString("Total", ":", 10, 10, 10, 10, 10, 10, 10, 10, 10, 10))

	slice := []int{10, 20, 30}
	total = sumAll(slice...)
	fmt.Println(total)
}