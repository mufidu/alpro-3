package main

import "fmt"

// Buat algoritma yang dapat membaca sebuah nilai uang (rupiah) yang harus dibayar dan lembar
// sepuluh ribuan yang disediakan. Algoritma tersebut akan menampilkan jumlah sepuluh ribuan paling
// sedikit yang dapat diperlukan untuk membayar nilai tersebut dan informasi apakah uang yang
// disediakan cukup atau tidak. Kalau cukup, sebutkan berapa kembaliannya.

// Contoh interaksi yang diminta adalah sebagai berikut (teks bergaris bawah adalah input):
// Nilai uang (Rp): 28500
// Lembar Rp.10000,- yang diberikan: 5
// Lembar Rp.10000,- yang digunakan: 3
// Cukup? true
// Kembalian: 21500
// Contoh interaksi lain adalah sebagai berikut (teks bergaris bawah adalah input):
// Nilai uang: 540000
// Lembar Rp.10000,- yang diberikan: 45
// Lembar Rp.10000,- yang digunakan: 54
// Cukup? false

// Keterangan:
// Contoh pertama, diperlukan 3 lembar sepuluh ribuan, sehingga 5 lembar sangat mencukupi.
// Contoh kedua, diperlukan 54 lembar sepuluh ribuan, sehingga 45 lembar sangat kurang.

func main() {
	var nominal, lembarSepuluhRibuDiberikan, lembarSepuluhRibuDigunakan, kembalian int
	var cukup bool

	fmt.Print("Nilai uang (Rp): ")
	fmt.Scanln(&nominal)
	fmt.Print("Lembar Rp.10000,- yang diberikan: ")
	fmt.Scanln(&lembarSepuluhRibuDiberikan) 
	
	if nominal % 10000 == 0 {
		lembarSepuluhRibuDigunakan = nominal / 10000
	} else {
		lembarSepuluhRibuDigunakan = nominal / 10000 + 1
	}

	fmt.Printf("Lembar Rp.10000,- yang digunakan: %d\n", lembarSepuluhRibuDigunakan)

	cukup = lembarSepuluhRibuDigunakan <= lembarSepuluhRibuDiberikan

	fmt.Println("Cukup?", cukup)

	if cukup {
		kembalian = lembarSepuluhRibuDiberikan * 10000 - nominal
		fmt.Println("Kembalian:", kembalian)
	}
}
