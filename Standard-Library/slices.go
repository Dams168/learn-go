package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"Alice", "Bob", "Charlie", "Diana", "Eve"}
	values := []int{10, 20, 30, 40, 50}

	// Menggunakan slices.Min untuk mendapatkan nilai minimum dari slice
	fmt.Println(slices.Min(names))
	fmt.Println(slices.Min(values))

	// Menggunakan slices.Max untuk mendapatkan nilai maksimum dari slice
	fmt.Println(slices.Max(names))
	fmt.Println(slices.Max(values))

	// Menggunakan slices.Contains untuk memeriksa apakah slice mengandung elemen tertentu
	fmt.Println(slices.Contains(names, "Charlie"))
	fmt.Println(slices.Contains(values, 25))

	// Menggunakan slices.Index untuk mendapatkan indeks dari elemen tertentu dalam slice
	fmt.Println(slices.Index(names, "Diana"))
	fmt.Println(slices.Index(values, 30))
}
