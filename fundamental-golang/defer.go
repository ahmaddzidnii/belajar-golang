package main

import "fmt"

func loging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication() {
	defer loging()
	fmt.Println("Run Application")
}
func main() {
	runApplication()
}
