package main

import "fmt"

func goodbye(name string) string {
	return "Goodbye " + name
}

func main() {
	gb := goodbye
	result := gb("Muhammad Edo Fadilah")
	fmt.Println(result)
}
