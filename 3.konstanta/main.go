package main

import "fmt"

func main() {
	// Deklarasi Konstanta Biasa
	const pi float64 = 3.14159
	const nama string = "Golang"

	// Deklari Konstanta Multiple
	const (
		negara  = "Indonesia"
		tahun   = 2025
		version = 1.0
	)

	fmt.Println("Pi:", pi)
	fmt.Println("Nama:", nama)
	fmt.Println("Negara:", nama, "Tahun:", tahun, "Versi:", version)

}
