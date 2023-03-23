package main

import "fmt"

func main() {
	for x := 0; x <= 10000; x++ {
		fmt.Println(x)
	}
	names := [...]string{"Muhammad", "Edo", "Fadilah"}
	for y := 0; y < len(names); y++ {
		fmt.Println(names[y])
	}
}
