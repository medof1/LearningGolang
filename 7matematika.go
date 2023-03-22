package main
import "fmt"

func main(){
	const(
		a = 1
		b = 2
	)
	var c = a+b
	fmt.Println(c)

	var i = 10
	i += 10
	fmt.Println(i)
	i = i +10
	fmt.Println(i)
	i++
	fmt.Println(i)
}