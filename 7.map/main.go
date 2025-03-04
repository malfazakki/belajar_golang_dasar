package main

import "fmt"

func main() {
	/*
		📌 1. Apa itu Map?

		Map di Go adalah struktur data key-value, mirip dengan Object di JavaScript atau Dictionary di Python.

		📝 Ciri-ciri Map di Go:
		✅ Menyimpan data dalam bentuk pasangan key-value
		✅ Key harus unik (tidak bisa ada duplikat)
		✅ Key bisa bertipe string, int, atau tipe lainnya
		✅ Value bisa bertipe apa saja
	*/

	// Deklarasi map dengan key string dan value int || map[keyType]valueType{}
	var nilai map[string]int = map[string]int{
		"Alice": 90,
		"Bob":   85,
	}
	fmt.Println("Data Nilai Map:", nilai)

	// Deklarasi map menggunakan make()
	nilai2 := make(map[string]int)

	// Menambahkan data kedalam map
	nilai2["Alice"] = 90
	nilai2["Bob"] = 85
	fmt.Println("Data Nilai Make:", nilai)

	// Mengubah nilai dalam map
	nilai2["Bob"] = 95
	fmt.Println("Data Nilai setelah diubah:", nilai2)

	/*
		* Menghapus elemen dari map | delete(map, key)
			- Menghapus data dengan key "Bob"
	*/
	delete(nilai2, "Bob")
	fmt.Println("Data Nilai setelah dihapus:", nilai2)

	/*
		* Mengecek key dalam map | value, exists := map[key]
			- Cek apakah key  ada dalam map
	*/
	value, exists := nilai2["Bob"]

	if exists {
		fmt.Println("Nilai Bob Adalah:", value)
	} else {
		fmt.Println("Data tidak ditemukan")
	}

	/*
		📝 Penjelasan:
		- Jika Charlie ada → exists = true, value berisi nilai.
		- Jika Charlie tidak ada → exists = false, value default (0 untuk int).
	*/
}
