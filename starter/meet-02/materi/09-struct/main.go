package main

import "fmt"

type BLKPP struct {
	Kelompok string
	ID int
	Nama string
}

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
	// fungsi baca
	//for ... {
	//	studentBlkpp = BLKPP{
	//		Kelompok: row[i].Kelompok
	//		ID: row[i].ID
	//		Nama: row[i].Nama
	//	}
	// }

	//mahasiswa := Mahasiswa{
	//	Nama:  "Andi",
	//	Umur:  20,
	//	Nilai: 82,
	//}

	blkpp := BLKPP{
		Kelompok: "Kelompok 1",
		ID: 1,
		Nama: "Budi",
	}

	fmt.Println("Nama:", blkpp.Nama)
	fmt.Println("ID:", blkpp.ID)
	fmt.Println("Kelompok:", blkpp.Kelompok)

	daftarMahasiswa := []Mahasiswa{
		{Nama: "Budi", Umur: 21, Nilai: 70},
		{Nama: "Siti", Umur: 20, Nilai: 90},
	}

	for _, data := range daftarMahasiswa {
		fmt.Println("\nNama", data.Nama)
		fmt.Println("Umur", data.Umur)
		fmt.Println("Status Lulus:", statusLulus(data.Nilai))
	}
}
