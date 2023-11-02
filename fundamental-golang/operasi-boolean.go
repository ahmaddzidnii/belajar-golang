package main

import "fmt"

func main() {
	var nilai = 90
	var presensi = 80

	var lulusNilai bool = nilai > 80
	var lulusPresensi bool = presensi > 80

	var lulus = lulusNilai && lulusPresensi // ini operasi boolean

	fmt.Println(lulus)
}
