package main

import "fmt"

/*
   Fungsi new() di Go
   Fungsi new() adalah built-in function di Go yang digunakan untuk mengalokasikan memori untuk sebuah tipe data tertentu
   dan mengembalikan pointer yang menunjuk ke lokasi memori tersebut.

   Fungsi new() menerima satu argumen yaitu tipe data yang ingin dialokasikan memorinya
   dan mengembalikan pointer ke tipe data tersebut dengan nilai awal (zero value) dari tipe data tersebut.

*/

type address struct {
	city, province, country string
}

func main() {
	alamat1 := new(address) // mengalokasikan memori untuk tipe data address dan mengembalikan pointer ke alamat1
	alamat2 := alamat1

	alamat2.city = "Cianjur"

	fmt.Println(alamat1)
	fmt.Println(alamat2)
}
