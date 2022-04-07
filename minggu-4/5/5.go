package main

import "fmt"

// Deret Fibonacci pertama kali ditemukan oleh matematikawan India pada abad pertengahan. Deret ini
// dimulai dengan 0 dan 1 sebagai nilai suku ke-0 dan suku ke-1. Apabila suku ke-i adalah Si, maka nilai
// suku untuk i > 2 adalah Si=Si-1 + Si-2.
// Contoh sepuluh suku pertama dari deret Fibonacci: 0 1 1 2 3 5 8 13 21 34 55.
// Nilai Suku ke-n dapat pula dihitung secara cukup akurat dengan menggunakan persamaan berikut:

// Buatlah algoritma untuk menghitung nilai suku ke-i dengan menggunakan formula di atas.
// Input adalah bilangan integer i, dan output adalah nilai Si. Proses berhenti apabila input adalah
// bilangan negatif.
// Catatan: Gunakan konstanta 2.236 sebagai pengganti nilai √5.
// Contoh input:
// 10
// 0
// -1
// Contoh output:
// 55
// 0
// Selesai

func main() {
	var i int
	for {
		fmt.Scanln(&i)
		if i == -1 {
			fmt.Println("Selesai")
			break
		} else {
			fmt.Println(fibonacci(i))
		}
	}	
}

func fibonacci(i int) int {
	if i == 0 {
		return 0
	} else if i == 1 {
		return 1
	} else {
		return fibonacci(i-1) + fibonacci(i-2)
	}
}