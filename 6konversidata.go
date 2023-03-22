package main
import "fmt"

func main(){
	const nilai64 int64 = 100000
	var nilai32 int32 = int32(nilai64)
	fmt.Println(nilai32)

	var name = "Edo"
	var e = name[0]
	fmt.Println(e)
	var estring = string(e)
	fmt.Println(estring)
}