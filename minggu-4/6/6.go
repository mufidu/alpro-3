package main

import "fmt"

// Seorang pelatih panahan akan mendata kekonsistenan atletnya dalam memanah dengan tepat sasaran.
// Sasaran panah terdiri dari 4 area lingkaran, mulai dari lingkaran terluar hingga lingkaran terdalam.
// Adapun skor per lingkaran dari yang terluar hingga terdalam adalah 1, 4, 7, 10. Di luar semua
// lingkaran tersebut, skor 0.
// Skema latihan dilakukan dalam sejumlah set tertentu, dimana 1 set latihan terdiri dari 3 kali
// kesempatan memanah. Pelatih akan mencatat skor yang dicapai sampai atlet berhasil mencapai target
// latihan, yaitu mampu memanah pada lingkaran terdalam dalam 1 set atau sudah mengumpulkan skor
// minimal 70. Perhatikan contoh-contoh di bawah ini (yang bergaris bawah adalah input)
// Contoh 1 :
// 10 10 10
// SELESAI
// Tercapai pada set ke-1
// Contoh 2 :
// 10 4 7
// 10 10 10
// SELESAI
// Tercapai pada set ke-2
// Contoh 3 :
// 10 4 4
// 10 7 10
// 0 10 7
// 4 1 10
// SELESAI
// Tercapai pada set ke-4
// Buatlah algoritma untuk membantu pelatih mencatat skor atletnya sampai target terpenuhi.

func main() {
	var skor1, skor2, skor3 int
	var totalSet, totalSkor int

	totalSkor = 0

	for totalSkor < 70 {
		fmt.Print("Masukkan 3 skor: ")
		fmt.Scanln(&skor1, &skor2, &skor3)
		totalSet += 1
		totalSkor += skor1 + skor2 + skor3

		if skor1 == 10 && skor2 == 10 && skor3 == 10 {
			break
		}
	}

	fmt.Printf("Tercapai pada set ke-%d\n", totalSet)
}