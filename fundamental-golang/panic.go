package main

import (
	"fmt"
)

func endApp() {
	fmt.Println("Aplikasi selesai")
}

func runApp(error bool) {
	if error {
		panic("APLIKASI ERROR")
	}
}
func main() {
	defer endApp()
	runApp(true)
}
