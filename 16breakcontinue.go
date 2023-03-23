package main

import "fmt"

func main() {
	nilai := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 9, 3, 3, 2, 2, 3, 4, 3, 2, 12, 5, 3, 53, 4, 321}
	for i := 1; i < len(nilai); i++ {
		for j := 0; j < len(nilai); j++ {
			if nilai[j] > nilai[i] {
				nilai[j], nilai[i] = nilai[i], nilai[j]
			} else {
				continue
			}
		}
	}
	fmt.Println(nilai)
}
