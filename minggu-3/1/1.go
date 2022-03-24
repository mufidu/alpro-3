package main

import "fmt"

// Buatlah sebuah program yang menerima inputan 2 angka,
// lalu outputkan permutasi dan kombinasi kedua angka tersebut.
// Buat 3 function untuk menghitung: faktorial, permutasi, dan kombinasi.

func main() {
	var a, b int
	fmt.Print("Masukkan angka pertama: ")
	fmt.Scanln(&a)
	fmt.Print("Masukkan angka kedua: ")
	fmt.Scanln(&b)

	fmt.Println("Permutasi: ", permutasi(a, b))
	fmt.Println("Kombinasi: ", kombinasi(a, b))
}

func faktorial(a int) int {
	// if a == 0 {
	// 	return 1
	// } else {
	// 	return a * faktorial(a-1)
	// }

	// Non recursive
	var f int
	f = 1
		
	for i := 1; i <= a; i++ {
		f *= i
	}
	
	return f
}

func permutasi(a int, b int) int {
	return faktorial(a) / faktorial(a-b)
}

func kombinasi(a int, b int) int {
	return faktorial(a) / (faktorial(a-b) * faktorial(b))
}