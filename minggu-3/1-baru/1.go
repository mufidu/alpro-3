package main

import "fmt"

// Buatlah sebuah program yang menerima inputan 2 angka,
// lalu outputkan permutasi dan kombinasi kedua angka tersebut.
// Buat 3 function untuk menghitung: faktorial, permutasi, dan kombinasi.

func main() {
	var n, r int

	fmt.Print("Masukkan angka pertama: ")
	fmt.Scanln(&n)
	fmt.Print("Masukkan angka kedua: ")
	fmt.Scanln(&r)

	fmt.Println("Permutasi:", permutasi(n, r))
	fmt.Println("Kombinasi:", kombinasi(n, r))
}

func faktorial(x int) int {
	var hasil int
	hasil = 1
	
	for i := 1; i <= x; i++ {
		hasil = hasil * i
	}

	return hasil
}

func permutasi(n, r int) int {
	var hasil int

	hasil = faktorial(n) / faktorial(n-r)
	
	return hasil
}

func kombinasi(n, r int) int {
	return faktorial(n) / (faktorial(r)*faktorial(n-r))
}