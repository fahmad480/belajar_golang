package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	faraaz := Man{"Faraaz"}
	faraaz.Married()

	fmt.Println(faraaz.Name)
}
