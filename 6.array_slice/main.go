package main

import "fmt"

func main() {
	/*
		Pengertian Array & Slice
			- Array: Struktur data dengan ukuran tetap (fixed size).
			- Slice: Struktur data berbasis array tetapi memiliki ukuran yang dinamis.
	*/

	// 1. Deklarasi Array (Ukuran Tetap)
	var angkaArray [5]int = [5]int{10, 20, 30, 40, 50}
	fmt.Println("Array:", angkaArray)
	fmt.Println("Index Ke-2:", angkaArray[2])

	// 2. Deklarasi Slice (Ukuran Dinamis)
	angkaSlice := []int{10, 20, 30}
	fmt.Println("Slice Awal:", angkaSlice)

	// 3. Modifikasi Elemen
	angkaArray[2] = 99
	angkaSlice[1] = 88
	fmt.Println("Array setelah modifikasi:", angkaArray)
	fmt.Println("Slice setelah modifikasi:", angkaSlice)

	// 4. Menambahkan Elemen ke Slice
	angkaSlice = append(angkaSlice, 40, 50)
	fmt.Println("Slice setelah Append:", angkaSlice)

	// 5. Menghapus Elemen dari Slice (Hapus Index ke-1)
	angkaSlice = append(angkaSlice[:2], angkaSlice[3:]...)
	fmt.Println("Slice setelah menghapus index ke-1:", angkaSlice)

	// 6. Menggabungkan Dua Slice
	sliceA := []int{1, 2, 3}
	sliceB := []int{4, 5, 6}
	sliceGabung := append(sliceA, sliceB...)
	fmt.Println("Hasil Gabung:", sliceGabung)

	// 7. Memotong Slice (Slicing)
	angka := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("Slice dari awal index 1-3:", angka[1:2])       // [20, 30, 40]
	fmt.Println("Slice dari awal sampai index 2:", angka[:2])   // [10, 20, 30]
	fmt.Println("Slice dari index 2 sampai akhir: ", angka[2:]) // [30, 40, 50]
}
