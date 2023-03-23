package main

import "fmt"

// Struct itu kayak class

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHi(name string) {
	fmt.Println("Hello", name, "My name is", customer.Name)
}

func main() {
	var edo Customer
	edo.Name = "Muhammad Edo Fadilah"
	edo.Address = "Bandung"
	edo.Age = 25

	fmt.Println(edo)
	fmt.Println(edo.Name)

	udin := Customer{
		Name:    "Udin God",
		Address: "Cibinong",
		Age:     100,
	}
	fmt.Println(udin)
	udin.sayHi("udin")

}
