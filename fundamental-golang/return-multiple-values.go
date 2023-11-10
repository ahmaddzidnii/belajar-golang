package main

import "fmt"

func getFullName() (string, string, string) {
	return "Ahmad", "Zidni", "Hidayat"
}

func main() {
	firstName, _, lastName := getFullName()
	fmt.Println(firstName, lastName)
}
