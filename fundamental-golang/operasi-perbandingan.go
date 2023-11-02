package main

import "fmt"

func main() {

	//operasi perbandingan adalah menghasilkan boolean
	var nama1 = "Ahmad"
	var nama2 = "Ahmad"

	var hasil bool = nama1 == nama2
	fmt.Println("Apakah kedua nama sama?", hasil)

	var value1 = 100
	var value2 = 200

	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)

}
