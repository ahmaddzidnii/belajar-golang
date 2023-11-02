package main

import "fmt"

func main() {
	var nilai32 int32 = 100
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	// konversi string
	var name = "Ahmad Zidni Hidayat"
	var a = name[0]
	var aString = string(a)

	fmt.Println(name)
	fmt.Println(aString)
}
