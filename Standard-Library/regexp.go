package main

/*
Package regexp di Go
Digunakan untuk bekerja dengan ekspresi reguler (regular expressions).
Package ini menyediakan fungsi dan tipe data untuk mencocokkan, mencari, dan memanipulasi string berdasarkan pola tertentu yang didefinisikan oleh ekspresi reguler.
Ekspresi reguler sering digunakan dalam pemrosesan teks, validasi input, dan pencarian pola dalam string.
*/

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile(`d[a-z, A-Z, 1-10]m`)

	fmt.Println(regex.MatchString("dam"))
	fmt.Println(regex.MatchString("dim"))
	fmt.Println(regex.MatchString("d1m"))
	fmt.Println(regex.MatchString("dIm"))

	search := regex.FindAllString("dam dim d1m dIm deym dumm", 10)
	fmt.Println(search)
}
