package main

import "fmt"

// kasus faktorial

// menggunakan loop
// func factorial(n int) int {
// 	result := 1
// 	for i := n; i > 0; i-- {
// 		result *= i
// 	}
// 	return result
// }

// func main() {
// 	fmt.Println(factorial(5)) // Output: 120
// 	fmt.Println(factorial(0)) // Output: 1
// }

//menggunakan recursive function

func factorialRecursive(n int) int {
	if n == 0 || n == 1 {
		return 1
	} else {
		return n * factorialRecursive(n-1)
	}
}

func main() {
	fmt.Println(factorialRecursive(5)) // Output: 120
	fmt.Println(factorialRecursive(0)) // Output: 1
}
