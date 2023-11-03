// slice aadalah potongan data array

package main

import (
	"fmt"
)

func main() {
	var month = [...]string{
		"januari",
		"februari",
		"maret",
		"april",
		"mei",
		"juni",
		"juli",
		"agustus",
		"september",
		"oktober",
		"november",
		"desember",
	}

	slice1 := month[4:7]
	fmt.Println(slice1) //output [mei juni juli]

	// Function slice

	//panjang slice
	fmt.Println(len(slice1)) // output 3

	// kapasitas slice
	fmt.Println(cap(slice1)) // output 8

	var slice2 = month[10:]
	fmt.Println(slice2) // output [november desember]

	var slice3 = append(slice2, "Ahmad")
	fmt.Println(slice3) // output [november desember Ahmad]

	slice3[1] = "Bukan Desember"
	fmt.Println(slice3) // output [november Bukan Desember Ahmad]

	fmt.Println(month) // output [januari februari maret april mei juni juli agustus september oktober november desember]

	// make slice

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Ahmad"
	newSlice[1] = "Zidni"

	fmt.Println(newSlice)      // output [Ahmad Zidni]
	fmt.Println(len(newSlice)) // output 2
	fmt.Println(cap(newSlice)) // output 5

	// copy slice

	// kpasitas harus sama agar data tidak terpotong
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice) // output [Ahmad Zidni]

	// Array vs Slice

	iniArray := [5]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray) // output [1 2 3 4 5]
	fmt.Println(iniSlice) // output [1 2 3 4 5]
}
