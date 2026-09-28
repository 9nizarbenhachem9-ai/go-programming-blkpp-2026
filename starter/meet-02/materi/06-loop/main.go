package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		fmt.Println("Perulangan ke-", i)
	}

	angka := []int{10, 20, 30}
	for index, value := range angka {
		fmt.Printf("Index %d berisi %d\n", index, value)
	}
}
