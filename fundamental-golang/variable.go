package main

import "fmt"

func main() {

	var name string
	name = "Ahmad"
	fmt.Println(name)
	name = "Ahmad Zidni"
	fmt.Println(name)

	// Default int itu 32 bit atau 64bit tergantung os
	var umur = 18
	fmt.Println(umur)

	// kalau mau spesifik
	var angka int8 = 10
	fmt.Println(angka)

	// var tidak wajib  namun perlu menggunakan kata kunci := saat menginisialisasikan data tersebut
	NIM := 77
	fmt.Println(NIM)

	// deklarasi multiple variable
	var (
		firstName = "Ahmad"
		lastName  = "Zidni"
		age       = 18
	)

	fmt.Println(firstName, lastName, age)

	// constant kayak const
	const salah bool = false
	const tetap = "Nurr.."

	// multiple constant
	const (
		firstName2 = "Ahmad"
		lastName2  = "Zidni"
		age2       = 18
	)
}
