package main

import "fmt"

func main() {
	var x, y int

	fmt.Print("Masukkan angka pertama: ")
	fmt.Scanln(&x)
	fmt.Print("Masukkan angka kedua: ")
	fmt.Scanln(&y)

	fmt.Println("Hasil function:", function(x, y))
}

func function(x, y int) int {
	var hasil int

	if y % x == 0 {
		hasil = 2*x + y
	} else if x % y == 0 {
		hasil = 2*y + x
	} else {
		hasil = 0
	}

	return hasil
}