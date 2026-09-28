package main

import (
	"fmt"
)

func bagi(a int, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("pembagi tidak boleh 0")
	}

	return a / b, nil
}

func main() {
	hasil, err := bagi(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("10 / 2 =", hasil)

	hasil, err = bagi(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("10 / 0 =", hasil)
}
