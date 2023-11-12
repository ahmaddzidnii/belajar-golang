package main

import "fmt"

type Adrress struct {
	City, Province, Country string
}

func main() {
	address1 := Adrress{
		City:     "Jakarta",
		Province: "DKI Jakarta",
		Country:  "Indonesia",
	} // Address

	address2 := &address1 // pointer Address

	address2.City = "Bandung"

	fmt.Println(address1)
	fmt.Println(address2)
}
