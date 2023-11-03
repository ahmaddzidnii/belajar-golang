package main

import "fmt"

func main() {
	person := map[string]string{
		"nama":      "Ahmad Zidni Hidayat",
		"umur":      "18",
		"pekerjaan": "Mahasiswa",
	}

	fmt.Println(person)              //output map[nama:Ahmad Zidni Hidayat umur:18 pekerjaan:Mahasiswa]
	fmt.Println(person["nama"])      //output Ahmad Zidni Hidayat
	fmt.Println(person["umur"])      //output 18
	fmt.Println(person["pekerjaan"]) //output Mahasiswa

	// function map

	var books = make(map[string]string)
	books["title"] = "2 jam menjadi sepuh react"
	books["author"] = "Ahmad Zidni Hidayat"
	books["year"] = "Ini salah"
	fmt.Println(books) //output map[author:Ahmad Zidni Hidayat title:2 jam menjadi sepuh react year:Ini salah]
	delete(books, "year")
	fmt.Println(books) //output map[author:Ahmad Zidni Hidayat title:2 jam menjadi sepuh react]
}
