package main

import "fmt"

func warga() (nomor int, nama, pekerjaan string) {
	nomor = 1023
	nama = "Muhammad Edp Fadollah"
	pekerjaan = "Hikikomori"
	return
}

func main() {
	fmt.Println(warga())
}
