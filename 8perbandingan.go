package main
import "fmt"

func main(){
	const(
		a = "M"
		b = "A"
	)
	var result bool = a < b
	fmt.Println(result) //false karena M lebih gede
}