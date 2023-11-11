package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {

	// cara pertama
	var ahmad Customer

	ahmad.Name = "Ahmad"
	ahmad.Address = "Jakarta"
	ahmad.Age = 19

	fmt.Println(ahmad)

	// cara kedua
	zidni := Customer{
		Name:    "Zidni",
		Address: "Jakarta",
		Age:     19,
	}

	fmt.Println(zidni)

	// cara ketiga
	hidayat := Customer{"Hidayat", "Jakarta", 19}
	fmt.Println(hidayat)
}
