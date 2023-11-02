package main

import "fmt"

func main() {
	var a = 10
	var b = 10
	var c = a + b
	var d = a - b
	var e = a * b
	var f = a / b
	var g = a % b
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)

	// augmented assignment
	a += 5 // a = a + 5
	b -= 3 // b = b - 3
	c *= 2 // c = c * 2
	d /= 2 // d = d / 2
	e %= 4 // e = e % 4

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

	// unary operator
	var angka = 100
	angka++ // angka = angka + 1
	angka-- // angka = angka - 1
	fmt.Println(angka)

	// negative
	var negatif = -100
	fmt.Println(negatif)

}
