package main

import "fmt"

func endApp() {
	fmt.Println("End App")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Error")
	}
	fmt.Println("App is running")
}

func main() {
	runApp(true)
	fmt.Println("App is done")
	fmt.Println("-------------------------------------------")
	runApp(false)
}