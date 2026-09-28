package main

import "fmt"

func main() {
	namaBarang := []string{"Buku", "Pulpen", "Tas"}
	namaBarang = append(namaBarang, "Penghapus")

	for _, barang := range namaBarang {
		fmt.Println("Barang:", barang)
	}

	stok := map[string]int{
		"Buku":   10,
		"Pulpen": 25,
		"Tas":    5,
	}

	fmt.Println("Stok Buku:", stok["Buku"])
	stok["Pulpen"] = 20
	fmt.Println("Stok Pulpen:", stok["Pulpen"])
}
