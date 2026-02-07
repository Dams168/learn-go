package main

/*
package csv di Go
Digunakan untuk membaca dan menulis data dalam format CSV (Comma-Separated Values).
Package ini menyediakan fungsi untuk mengurai data CSV dari string atau file, serta menulis data ke format CSV.
Format CSV sering digunakan untuk pertukaran data antar aplikasi, terutama dalam konteks spreadsheet dan basis data.
*/

import (
	"encoding/csv"
	"fmt"
	"strings"
)

func main() {
	// Contoh data CSV dalam bentuk string
	csvString := `John,Doe,30\n
Abdul,Smith,25\n
Jane,Nur,40`

	// Membuat reader CSV dari string
	// csv.NewReader membutuhkan io.Reader sebagai inputnya
	// kita bisa menggunakan strings.NewReader untuk mengubah string menjadi io.Reader
	// io reader adalah interface yang digunakan untuk membaca data secara berurutan

	reader := csv.NewReader(strings.NewReader(csvString))

	// Perulangan untuk membaca semua record dalam data CSV
	// perulangan ini dinamakan loop infinito, yang akan terus berjalan sampai kita menghentikannya secara eksplisit
	for {
		// Membaca setiap record (baris) dari data CSV
		// record adalah slice of string yang berisi nilai-nilai dari setiap kolom dalam baris tersebut
		// jika sudah tidak ada data lagi, maka akan mengembalikan error io.EOF
		// kita bisa menggunakan error ini untuk menghentikan perulangan
		record, err := reader.Read()
		if err != nil {
			break
		}

		fmt.Println(record)

	}

}
