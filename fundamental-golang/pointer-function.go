package main

import "fmt"

type Adrress struct {
	City, Province, Country string
}

func changeCountryToIndonesia(address *Adrress) {
	address.Country = "Indonesia"
}

func main() {
	address := Adrress{}
	changeCountryToIndonesia(&address)
	fmt.Println(address)
}
