package main
import "fmt"

//perbedaan sama array, slice jumlahnya bisa berubah
func main() {
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}
	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))
	fmt.Println(len(months))
	fmt.Println(cap(months))

	months[4] = "BulanSABIT"
	fmt.Println(slice1)

	var slice2 = months[10:]
	fmt.Println(slice2)
	fmt.Println(cap(slice2))

	slice2 = append(slice2, "Udin")
	fmt.Println(slice2)
	fmt.Println(cap(slice2))
	fmt.Println(months)
	slice1 = append(slice1, "Ucok")
	fmt.Println(slice1)
	fmt.Println(months)
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))

	newSlice := make([]string, 2, 5)

	newSlice[0] = "Udin"
	newSlice[1] = "Pakabs"

	fmt.Println(len(newSlice))

	copySlice := make([]string, len(slice1), cap(slice1))
	copy(copySlice, slice1)
	fmt.Println(copySlice)
}