package main

import "fmt"

func main() {
	name := "Ahmad"

	if name == "Ahmad" {
		fmt.Println("Hello Ahmad")
	} else if name == "Zidni" {
		fmt.Println("Hello Zidni")
	} else if name == "Hidayat" {
		fmt.Println("Hello Hidayat")
	} else {
		fmt.Println("Hi boleh kenalan?")
	}

	// if dengan short statement
	if age := 18; age >= 18 {
		fmt.Println("You are an adult")
	} else {
		fmt.Println("You are not an adult")
	}

}
