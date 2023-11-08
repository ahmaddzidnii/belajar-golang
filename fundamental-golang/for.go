package main

import "fmt"

func main() {
	counter := 1
	for counter <= 10 {
		fmt.Println("Perulangan ke", counter)
		counter++
	}

	for counter := 1; counter < 10; counter++ {
		fmt.Println("Perulangan ke", counter)
	}

	names := []string{"Ahmad", "Zidni", "Hidayat"} //tipe slice

	for index, name := range names {
		fmt.Println("Index", index, "Name", name)
	}

	for _, name := range names {
		fmt.Println("Name", name) // jika tidak butuh index bisa dituliss begini
	}

}
