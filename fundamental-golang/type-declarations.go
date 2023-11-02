package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKtpAhmad NoKTP = "1234567890123456"
	var isMarried Married = false
	
	fmt.Println(noKtpAhmad)
	fmt.Println(isMarried)
}
