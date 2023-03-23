package main

import "fmt"

func endApp() {
	fmt.Println("App Selesai")
	message := recover()
	if message != nil {
		fmt.Println(message)
	}
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("App Error")
	}
	fmt.Println("App berjalan")
}

func main() {
	runApp(true)
}
