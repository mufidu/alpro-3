package main

import "fmt"

// Fakultas Informatika sedang mendata para lulusan jenjang sarjana. Informasi yang dicatat ada 3:
// masa studi dalam semester, indeks prestasi kumulatif (IPK), dan publikasi. Lulusan dinyatakan
// memiliki predikat cumlaude jika masa studi maksimal 8 semester, IPK minimal 3.51, dan memiliki
// publikasi. Fakultas ingin mengetahui banyaknya lulusan cumlaude dan rerata IPK lulusan.

// Buatlah algoritma yang menerima sejumlah data lulusan: masa studi m (6 < m < 14), ipk (2.0 < ipk <
// 4.0), dan publikasi [false, true]. Input diakhiri dengan MARK [0, 0, false]. Algoritma harus
// menampilkan banyaknya lulusan cumlaude dan rerata IPK lulusan.

// Contoh input/output (tulisan bergaris bawah menyatakan masukan, lainnya merupakan keluaran):
// 2.0 12 false
// 3.0 10 false
// 3.6 7 true
// 3.8 8 true
// 3.7 7 false
// 3.9 8 true
// 0 0 false
// Lulusan cumlaude : 3
// Rerata IPK lulusan: 3.333

func main() {
	var ipk float64
	var masaStudi int
	var publikasi bool

	var lulusanCumlaude int
	var jumlahIPK float64
	var jumlahLulusan int

	for {
		fmt.Scanln(&ipk, &masaStudi, &publikasi)

		if ipk < 2 || ipk > 4 || masaStudi < 6 || masaStudi > 14 {
			continue
		} else if ipk == 0 && masaStudi == 0 && publikasi == false {
			break
		} else if ipk >= 3.51 && masaStudi <= 8 && publikasi == true {
			lulusanCumlaude++
		}

		jumlahIPK += ipk
		jumlahLulusan++
	}

	fmt.Println("Lulusan cumlaude :", lulusanCumlaude)
	fmt.Println("Rerata IPK lulusan:", jumlahIPK/float64(jumlahLulusan))
}