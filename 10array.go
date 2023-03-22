package main
import "fmt"

func main() {
	var names = [3] string{
		"Muhammad",
		"Edo",
		"Fadilah",
	}
	for a:=0; a<len(names); a++{
		names[2] = "Fadilceut"
		fmt.Println(names[a])
	}
}