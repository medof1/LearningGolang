package main

import "fmt"

func mahasiswa(nim int, name string, nilai int) {
	fmt.Println("NIM: ", nim, "\nNama: ", name, "\nNilai: ", nilai)
}

func main() {
	mahasiswa(31625407, "Muhammad Edo Fadilah", 100)
}
