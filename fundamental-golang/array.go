package main

import "fmt"

func main() {
	var array [3]string
	array[0] = "Ahmad"
	array[1] = "Zidni"
	array[2] = "Hidayat"

	fmt.Println(array[0], array[1], array[2])

	var array2 = [3]string{"Ahmad",
		"Zidni",
		"Hidayat",
	}
	fmt.Println(array2[0], array2[1], array2[2])

	// Function Array
	var array3 = [2]int{
		11,
		22,
	}

	fmt.Println(len(array3))
}
