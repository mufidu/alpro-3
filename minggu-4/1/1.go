package main

import "fmt"

func main() {
	// Fungsi sumasi
	// Menghitung jumlah angka di interval a dan b

	var a, b, x int
	x = 0

	fmt.Print("a, b = ")
	fmt.Scanln(&a, &b)

	for i := a; i <= b; i++ {
		x += i
	}

	fmt.Println("nilai x =", x)
}