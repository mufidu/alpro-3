package main

import "fmt"

func main() {
	// Outputkan hasil dari 3/n - 3/(n+1) + 3/(n+2) - 3/(n+3) + ... + 3/m

	// Deklarasi dan inisialisasi variabel m dan n
	var m, n int;
	fmt.Scanln(&n);
	fmt.Scanln(&m);

	// Deklarasi dan inisialisasi variabel hasil
	var hasil float64;
	hasil = 3/float64(n);

	// j untuk selang seling +/-
	var j = 0;

	for i := n+1; i <= m; i++ {
		// gunakan j untuk menentukan +/-
		if j % 2 == 0 {
			hasil -= 3/float64(i);
		} else {
			hasil += 3/float64(i);
		}
		j++;
	}

	fmt.Printf("%.2f", hasil);
	fmt.Println();
}