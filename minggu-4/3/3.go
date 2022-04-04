package main

import "fmt"

// Salah satu syarat lulus Universitas Telkom adalah penguasaan bahasa Inggris dalam bentuk
// nilai English Proficiency Test (EPrT), Lulusan harus mempunyai nilai EPrT minimum
// 500 atau setidaknya telah mengulang EPrT sebanyak 3 kali.

// Buat algoritma untuk memproses data calon wisudawan yang telah memenuhi syarat EPrT.
// Masukan adalah nilai EPrT yang pernah diambil oleh N calon wisudawan, ada maksimum 3
// nilai EPrT untuk tiap calon. Hitung rata-rata nilai EPrT untuk seluruh wisudawan, dimana
// nilai yang diambil adalah nilai yang sudah mencapai minimum 500 atau nilai terakhir dari 3
// nilai yang diperoleh jika semuanya dibawah 500.

// Perhatikan contoh interaksi berikut (teks bergaris bawah adalah masukan. Pada contoh
// dibawah, nilai yang diperhitungkan untuk rerata adalah 520, 500, 510, 430, dan 610)
// Jumlah calon wisudawan: 5
// Calon 1 EPrT 1: 520
// Calon 2 EPrT 1: 470
// Calon 2 EPrT 2: 500
// Calon 3 EPrT 1: 440
// Calon 3 EPrT 2: 470
// Calon 3 EPrT 3: 510
// Calon 4 EPrT 1: 420
// Calon 4 EPrT 2: 445
// Calon 4 EPrT 3: 430
// Calon 5 EPrT 1: 610
// Rerata EPrT: 514.0

func main() {
	var jumlahCalon int
	var EPrT, totalEPrT int
	var rerataEPrT float64
	
	fmt.Print("Jumlah calon wisudawan: ")
	fmt.Scanln(&jumlahCalon)

	for i := 0; i < jumlahCalon; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print("Calon ", i+1, " EPrT ", j+1, ": ")
			fmt.Scanln(&EPrT)
			if EPrT >= 500 {
				break
			}
		}
		totalEPrT += EPrT
	}

	rerataEPrT = float64(totalEPrT) / float64(jumlahCalon)
	fmt.Println("Rerata EPrT:", rerataEPrT)
}
