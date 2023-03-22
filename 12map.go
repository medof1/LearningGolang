package main
import "fmt"

func main() {
	var person = map[string]string{
		"name": "Edo",
		"address": "Bandung",
	}

	person["title"] = "Pengangguran"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(len(person))
	delete(person, "title")
	fmt.Println(person)

	var book map[string]string = make(map[string]string)
	book["title"] = "Harry Potong"
	book["author"] = "me"

	fmt.Println(book)

}