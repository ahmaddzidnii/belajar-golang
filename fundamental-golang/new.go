package main

import "fmt"

type Adrress struct {
	City, Province, Country string
}

func main() {
	var alamat1 *Adrress = new(Adrress)
	var alamat2 = alamat1

	alamat2.Country = "Indonesia"

	fmt.Println(alamat1)
	fmt.Println(alamat2)
}
