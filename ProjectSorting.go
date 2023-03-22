package main
import "fmt"

func main() {
	var nilai= [...] int {123, 24, 351, 573, 43, 75, 65, 4, 3, 7, 8, 0, 2, 42, 2, 23, 35, 3, 54, 5, 31, 49, 12, 1}
	for i:= 1; i <len(nilai); i++{
		for j:= 0; j<len(nilai); j++{
			if nilai[j] > nilai[i]{
				nilai[j], nilai[i] = nilai[i], nilai[j]
			}
		}
	}
	fmt.Println(nilai)
}