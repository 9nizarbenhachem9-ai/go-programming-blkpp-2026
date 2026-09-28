package main

import "fmt"

func main() {
	nilai := 82

	if nilai >= 90 {
		fmt.Println("Grade A")
	} else if nilai >= 75 {
		fmt.Println("Grade B")
	} else if nilai >= 60 {
		fmt.Println("Grade C")
	} else {
		fmt.Println("Belum lulus")
	}

	hari := "senin"

	switch hari {
	case "sabtu", "minggu":
		fmt.Println("Akhir pekan")
	default:
		fmt.Println("Hari belajar")
	}
}
