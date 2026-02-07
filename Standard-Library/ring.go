package main

/*
	Package Ring Go
	Package ring ini berfungsi untuk mengimplementasikan struktur data lingkaran (circular list) di mana elemen terakhir terhubung kembali ke elemen pertama, sehingga membentuk sebuah lingkaran.
	Struktur data ini berguna dalam berbagai aplikasi seperti penjadwalan tugas, buffer melingkar, dan permainan.
	Pustaka ini menyediakan berbagai fungsi untuk membuat, mengelola, dan memanipulasi lingkaran, termasuk menambahkan, menghapus, dan mengiterasi elemen-elemen di dalamnya.


*/

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	data := ring.New(5)

	for i := 0; i < data.Len(); i++ {
		data.Value = "Value " + strconv.Itoa(i)
		data = data.Next()
	}

	// Menampilkan elemen-elemen dalam ring
	// Menggunakan metode Do untuk mengiterasi dan mencetak setiap elemen dalam ring
	// Fungsi Do menerima sebuah fungsi sebagai argumen yang akan dipanggil untuk setiap elemen dalam ring
	data.Do(func(x interface{}) {
		fmt.Println(x.(string))
	})
}
