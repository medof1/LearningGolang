package main
import "fmt"

func main(){
	const(
		a = true
		b = false
	)
	fmt.Println(a && b)
	fmt.Println(a || b)

	var ujian = 80
	var absensi = 80

	var lulusUjian = ujian >= 80
	var lulusAbsensi = absensi >= 80

	var lulus = lulusUjian && lulusAbsensi
	fmt.Println(lulus)
}