package main

import "fmt"

func getGoodBye(name string) string {
	return "Good Bye " + name
}

func main() {
	// function as value
	goodBye := getGoodBye
	fmt.Println(goodBye("Ahmad"))
}
