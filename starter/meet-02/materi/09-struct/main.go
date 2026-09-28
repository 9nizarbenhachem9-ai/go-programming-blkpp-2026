package main

import "fmt"

type Mahasiswa struct {
	Nama  string
	Umur  int
	Nilai int
}

func statusLulus(nilai int) string {
	if nilai >= 75 {
		return "Lulus"
	}

	return "Belum lulus"
}

func main() {
	mahasiswa := Mahasiswa{
		Nama:  "Andi",
		Umur:  20,
		Nilai: 82,
	}

	fmt.Println("Nama:", mahasiswa.Nama)
	fmt.Println("Umur:", mahasiswa.Umur)
	fmt.Println("Nilai:", mahasiswa.Nilai)
	fmt.Println("Status:", statusLulus(mahasiswa.Nilai))

	daftarMahasiswa := []Mahasiswa{
		{Nama: "Budi", Umur: 21, Nilai: 70},
		{Nama: "Siti", Umur: 20, Nilai: 90},
	}

	for _, data := range daftarMahasiswa {
		fmt.Printf("%s - %s\n", data.Nama, statusLulus(data.Nilai))
	}
}
