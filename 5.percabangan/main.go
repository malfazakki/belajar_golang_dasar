package main

import "fmt"

func main() {
	nilai := 85

	// 1. Percabangan If-Else
	if nilai >= 90 {
		fmt.Println("Grade: A")
	} else if nilai >= 80 {
		fmt.Println("Grade: B")
	} else if nilai >= 70 {
		fmt.Println("Grade: C")
	} else {
		fmt.Println("Grade: D")
	}

	// 2. Percabangan Switch Case
	hari := "Senin"

	switch hari {
	case "Senin":
		fmt.Println("Hari ini adalah awal minggu")
	case "Jum'at":
		fmt.Println("Hari ini adalah hari raya Ummat Islam")
	case "Ahad":
		fmt.Println("Hari ini libur!")
	default:
		fmt.Println("Hari Kerja")
	}

	// 3. Percabangan Switch dengan multiple case
	angka := 2

	switch angka {
	case 2, 4, 6, 8, 10:
		fmt.Println("Angka Genap")
	case 1, 3, 5, 7, 9:
		fmt.Println("Angka Ganjil")
	default:
		fmt.Println("Diluar Angka 1-10")
	}

	/*
		Penjelasan:
		- if-else → Digunakan untuk kondisi yang fleksibel dengan banyak syarat.
		- switch-case → Digunakan untuk membandingkan satu nilai dengan beberapa kemungkinan (lebih bersih dibanding if-else).
		- switch dengan multiple case → Bisa mengecek beberapa nilai dalam satu case.
	*/
}
