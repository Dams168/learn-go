package main

import "fmt"

/*
   Asterisk Operator (*)
   Asterisk (*) adalah operator yang digunakan untuk mendeklarasikan sebuah variable sebagai pointer
   atau untuk mengakses value dari sebuah pointer.

   Penggunaan Asterisk (*)
   1. Deklarasi Pointer
   2. Mengakses Value dari Pointer
   3. Mengubah Value melalui Pointer
   4. Pointer ke Struct

*/

type address struct {
	city, province, country string
}

func main() {
	var addr1 address = address{"Bandung", "Jawa Barat", "Indonesia"}
	var addr2 *address = &addr1

	// addr2.city = "Cianjur"

	fmt.Println(addr1) // data asli
	fmt.Println(addr2) // pointer yang mereferensikan addr1

	// Mengubah value melalui pointer, dengan menggunakan asterisk (*)
	// kita bisa mengubah value dari addr1 melalui addr2, karena addr2 adalah pointer yang mereferensikan alamat memori dari addr1
	*addr2 = address{"Cianjur", "Jawa Barat", "Indonesia"}

	fmt.Println(addr1) // data asli setelah diubah melalui pointer
	fmt.Println(addr2) // pointer yang mereferensikan addr1 setelah diubah
}
