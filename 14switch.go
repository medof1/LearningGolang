package main
import "fmt"

func main() {
	name := "endogan"

	switch name{
	case "edo":
		fmt.Println("Pakabs", name)
	default:
		fmt.Println("WHO YOU?", name)
	}
}