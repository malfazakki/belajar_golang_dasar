package main

import "fmt"

func main() {
	// Deklarasi dengan tipe data explisit
	var nama string = "Budi"
	var umur int = 25

	// Deklarasi dengan tipe data otomatis
	kota := "Jakarta"
	tinggi := 166

	// Deklarasi multiple
	var a, b, c int = 1, 2, 3

	fmt.Println("Nama:", nama, "|", "Umur:", umur)
	fmt.Println("Asal:", kota, "|", "Tinggi:", tinggi)
	fmt.Println("Nilai:", a, b, c)
}
