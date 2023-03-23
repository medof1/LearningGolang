package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	fmt.Println(sumAll(0, 1, 4, 23, 54, 34, 2, 1, 2, 43, 54, 544))
	num := []int{2, 32, 43, 34, 4, 4, 43, 43}
	fmt.Println(sumAll(num...))
}
