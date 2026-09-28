package main

import "fmt"

func main() {
	var nama string
	var umur int

	// fmt.Print("Masukkan nama: ")
	// fmt.Scanln(&nama)

	// fmt.Print("Masukkan umur: ")
	// fmt.Scanln(&umur)

	fmt.Print("Masukkan nama: ")
	_, err := fmt.Scanln(&nama)
	if err != nil {
		fmt.Println("Gagal membaca nama:", err)
		return
	}

	fmt.Print("Masukkan umur: ")
	_, err = fmt.Scanln(&umur)
	if err != nil {
		fmt.Println("Gagal membaca umur. Pastikan umum adalah angka!", err)
		return
	}

	fmt.Printf("Halo %s, umur kamu %d tahun.\n", nama, umur)
}
