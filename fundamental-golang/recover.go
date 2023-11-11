// recover menangkap data panic
package main

import (
	"fmt"
)

func endApp() {
	fmt.Println("Aplikasi selesai")
	message := recover()

	fmt.Println("Terjadi error", message)
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
