package main

/*
Package base64 di Go
Digunakan untuk encoding dan decoding data dalam format Base64.
Package ini menyediakan fungsi untuk mengonversi data biner menjadi string yang dapat dicetak dan sebaliknya, yang berguna untuk transmisi data melalui media yang hanya mendukung teks.
Base64 sering digunakan dalam konteks seperti penyimpanan data dalam format teks, pengiriman email, dan penyisipan data biner dalam dokumen teks.


*/

import (
	"encoding/base64"
	"fmt"
)

func main() {
	value := "Hello, World!"

	// Encoding
	// Dengan menggunakan StdEncoding untuk mengonversi string ke format Base64
	encoded := base64.StdEncoding.EncodeToString([]byte(value))
	fmt.Println(encoded)

	// Decoding
	// Mengonversi kembali string Base64 ke format aslinya
	// kenapa harus di-handle error? karena proses decoding bisa gagal jika string yang diberikan bukan format Base64 yang valid
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("Error decoding:", err)
		return
	}
	fmt.Println(string(decoded))
}
