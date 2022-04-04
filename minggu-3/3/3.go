package main

import "fmt"

// Buatlah program yang menerima input 3 bilangan,
// dan buat sebuah function yang mengoutputkan
// semua angka di interval 3 bilangan tadi.
// Contoh:
// Input: 1 3 6
// Output: 1 2 3 4 5 6

func main() {
	var a, b, c int
	fmt.Print("Masukkan angka pertama: ")
	fmt.Scanln(&a)
	fmt.Print("Masukkan angka kedua: ")
	fmt.Scanln(&b)
	fmt.Print("Masukkan angka ketiga: ")
	fmt.Scanln(&c)

	fmt.Println("Interval: ")
	interval(a, b, c)
}

func interval(a int, b int, c int) {
	for i := a; i <= b; i++ {
		fmt.Print(i, " ")
	}

	for i := b+1; i <= c; i++ {
		fmt.Print(i, " ")
	}

	fmt.Println()
}
