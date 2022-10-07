package main

import "fmt"

func endApp() {
	message := recover()
	if message != nil {
		fmt.Println("Error with message", message)
	}
	fmt.Println("End App")
}

func runApp(error bool) {
	defer endApp()
	fmt.Println("Run App")
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