package main

import "fmt"

func main() {
	// 1. Operator Aritmatika
	a, b := 10, 3
	fmt.Println("Penjumlahan:", a+b) // 13
	fmt.Println("Pengurangan:", a-b) // 7
	fmt.Println("Perkalian:", a*b)   // 30
	fmt.Println("Pembagian:", a/b)   // 3 (karena int)
	fmt.Println("Modulus:", a%b)     // 1 (sisa bagi)

	// 2. Operator Perbandingan
	fmt.Println("Apakah a > b?", a > b)   // true
	fmt.Println("Apakah a == b?", a == b) // false
	fmt.Println("Apakah a != b?", a != b) // true

	// 3. Operator Logika
	x, y := true, false
	fmt.Println("AND:", x && y) // false (karena salah satu false)
	fmt.Println("OR:", x || y)  // true (karena salah satu true)
	fmt.Println("NOT:", !x)     // false (kebalikan dari true)

	// 4. Operator Penugasan
	angka := 5
	angka += 2                                 // Sama dengan angka = angka + 2
	fmt.Println("Angka setelah += 2: ", angka) // 7
}
