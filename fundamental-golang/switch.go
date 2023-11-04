package main

import "fmt"

func main() {
	name := "Ahmad"

	switch name {
	case "Ahmad":
		fmt.Println("Hello Ahmad") // output Hello Ahmad

	case "Zidni":
		fmt.Println("Hello Zidni") // output Hello Zidni

	default:
		fmt.Println("Hi boleh kenalan?") // output Hi boleh kenalan?
	}

	// switch dengan short statement

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Terlalu Panjang") // output Terlalu Panjang
	case false:
		fmt.Println("Halo") // output Halo
	}

	// switch tanpa kondisi

	length := len(name)
	switch {
	case length > 15:
		fmt.Println("Terlalu Panjang") // output Terlalu Panjang
	case length > 10:
		fmt.Println("Panjang") // output Panjang
	default:
		fmt.Println("Halo") // output Halo
	}

}
