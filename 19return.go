package main

import "fmt"

func mahasiswa(nim int, name string, nilai int) (string, string) {
	fmt.Println("\nNIM: ", nim, "\nNama: ", name, "\nNilai: ", nilai)
	if nilai >= 80 {
		return "Lulus", ", Good job sir"
	} else {
		return "Drop Out", "Go Mulung now"
	}
}

func main() {
	result, _ := mahasiswa(31625407, "Muhammad Edo Fadilah", 10)
	fmt.Println(result)
	fmt.Println(mahasiswa(31625403, "Udin Junaedy", 100))
}
