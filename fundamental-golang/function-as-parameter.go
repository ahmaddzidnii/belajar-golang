package main

import "fmt"

// func sayHelloWithFilter(name string, filter func(string) string) {
// 	filteredParameter := filter(name)
// 	fmt.Println("Hello", filteredParameter)
// }

// func spamFilter(name string) string {
// 	if name == "Anjing" {
// 		return "..."
// 	} else {
// 		return name
// 	}
// }
// func main() {

// 	sayHelloWithFilter("Anjing", spamFilter) // output ...

// 	sayHelloWithFilter("Ahmad Zidni", spamFilter) // output Ahmad Zidni
// }

// function type declaration

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	filteredParameter := filter(name)
	fmt.Println("Hello", filteredParameter)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}
func main() {

	sayHelloWithFilter("Anjing", spamFilter) // output ...

	sayHelloWithFilter("Ahmad Zidni", spamFilter) // output Ahmad Zidni
}
