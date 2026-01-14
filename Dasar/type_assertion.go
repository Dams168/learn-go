package main

import "fmt"

/*
type assertion adalah cara untuk mengkonversi tipe data dari tipe interface kosong (any) ke tipe data yang diinginkan
dengan menggunakan sintaks => value.(TipeData)
jika tipe data sesuai, maka akan berhasil dan menghasilkan nilai dengan tipe data yang diinginkan
jika tidak sesuai, maka akan terjadi panic
untuk menghindari panic, kita bisa menggunakan bentuk kedua dari type assertion
yaitu dengan menambahkan variabel kedua yang bertipe boolean untuk mengecek apakah konversi berhasil atau tidak
contoh:
value, ok := result.(string)
if ok {
	// konversi berhasil
} else {
	// konversi gagal
}

kita juga bisa menggunakan type assertion dalam switch statement untuk mengecek tipe data dari sebuah nilai
dengan menggunakan sintaks => switch value := result.(type) { ... }

*/

func random() any {
	return "OK"
}

func main() {
	result := random()

	switch value := result.(type) {
	case string:
		fmt.Println("String:", value)
	case int:
		fmt.Println("Integer:", value)
	default:
		fmt.Println("Unknown Type", value)
	}
}
