package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHi(name string){
	fmt.Println("Hello", name, "My Name is", customer.Name)
}

func (a Customer) sayHuuu(){
	fmt.Println("Huuuuuu from", a.Name)
}

func main() {
	var faraaz Customer
	faraaz.Name = "Faraaz"
	faraaz.Address = "Indonesia"
	faraaz.Age = 30

	faraaz.sayHi("Joko")
	faraaz.sayHuuu()

	//fmt.Println(faraaz.Name)
	//fmt.Println(faraaz.Address)
	//fmt.Println(faraaz.Age)
	//
	//joko := Customer{
	//	Name:    "Joko",
	//	Address: "Cirebon",
	//	Age:     35,
	//}
	//fmt.Println(joko)
	//
	//budi := Customer{"Budi", "Jakarta", 35}
	//fmt.Println(budi)
}
