package main

import "fmt"

// catatan harus hati hati saat menggunakan type assertions karena jika salah tipe data maka akan error
func random() interface{} {
	return true
}
func main() {
	result := random()

	// manual
	// resultString := result.(string)
	// fmt.Println(resultString)

	// cara aman menggunakan type assertions
	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Integer", value)
	case bool:
		fmt.Println("Boolean", value)
	default:
		fmt.Println("Unknown", value)
	}

}
