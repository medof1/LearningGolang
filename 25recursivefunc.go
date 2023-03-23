package main

import "fmt"

func factorial(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

func factorialrecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialrecursive(value-1)
	}
}

func main() {
	loop := factorial(5)
	fmt.Println(loop)
	fmt.Println(factorialrecursive(5))
}
