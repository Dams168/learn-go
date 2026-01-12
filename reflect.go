package main

import (
	"fmt"
	"reflect"
)

// reflect adalah paket yang menyediakan kemampuan untuk memeriksa tipe data dan nilai pada runtime
// dengan menggunakan reflect, kita dapat memperoleh informasi tentang tipe data, nilai, dan struktur data secara dinamis
// reflect sering digunakan dalam situasi di mana tipe data tidak diketahui pada saat kompilasi
// atau ketika kita ingin membuat fungsi yang dapat bekerja dengan berbagai tipe data
func main() {
	number := 42

	reflectValue := reflect.ValueOf(number)

	fmt.Println("Tipe data:", reflectValue.Type())

	if reflectValue.Kind() == reflect.Int {
		fmt.Println("Nilai:", reflectValue.Int())
	}
}
