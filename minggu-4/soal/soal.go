package main

import "fmt"

// 2 Function 1 Procedure
// 1. nilai minimum dari 2 bilangan
// 2. nilai maksimum dari 2 bilangan
// 3. procedure menggunakan function min dan max

func main() {
	var a, b int
	fmt.Print("Masukkan angka pertama: ")
	fmt.Scanln(&a)
	fmt.Print("Masukkan angka kedua: ")
	fmt.Scanln(&b)

	fmt.Println("Nilai minimum: ", min(a, b))
	fmt.Println("Nilai maksimum: ", max(a, b))
}

func min(a, b int) int {
	var min int

	if a < b {
		min = a
	} else {
		min = b
	}

	return min
}

func max(a, b int) int {
	var max int

	if a > b {
		max = a
	} else {
		max = b
	}

	return max
}