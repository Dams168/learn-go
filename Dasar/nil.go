package main

import "fmt"

/*
nil adalah nilai default untuk tipe map,interface, pointer, slice, channel, dan function
jadi kita bisa mengembalikan nil jika tidak ada data
untuk menghemat penggunaan memori
karena map yang nil tidak mengalokasikan memori
dan tidak bisa digunakan sebelum diinisialisasi
jadi kita harus hati-hati saat menggunakannya
jika kita mencoba menambahkan elemen ke map yang nil, akan terjadi panic
jadi pastikan untuk memeriksa apakah map tersebut nil sebelum menggunakannya
atau inisialisasi terlebih dahulu sebelum digunakan
*/

func newMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	m := newMap("Yono")

	if m == nil {
		fmt.Println("Data Map masih Kosong")
	} else {
		fmt.Println("Data Map:", m["name"])
	}
}
