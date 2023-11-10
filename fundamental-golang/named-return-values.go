package main

import "fmt"

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Ahmad"
	middleName = "Zidni"
	lastName = "Hidayat"
	return firstName, middleName, lastName
}

func main() {
	firstName, middleName, lastName := getCompleteName()

	fmt.Println(firstName, middleName, lastName)
}
