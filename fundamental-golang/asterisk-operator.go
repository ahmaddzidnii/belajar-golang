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

	fmt.Println(address1) // output: {Jakarta DKI Jakarta Indonesia}
	fmt.Println(address2) // output: &{Bandung DKI Jakarta Indonesia}

	*address2 = Adrress{
		City:     "Surabaya",
		Province: "Jawa Timur",
		Country:  "Indonesia",
	} //* operator asterisk
	fmt.Println(address1) // output: {Surabaya Jawa Timur Indonesia}
	fmt.Println(address2) // output: &{Surabaya Jawa Timur Indonesia}

}
