package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	fmt.Println(filepath.Dir("hello/world.go"))              //mengembalikan direktori dari path yang diberikan
	fmt.Println(filepath.Base("hello/world.go"))             //mengembalikan nama file dari path yang diberikan
	fmt.Println(filepath.Ext("hello/world.go"))              //mengembalikan ekstensi file dari path yang diberikan
	fmt.Println(filepath.Join("hello", "world", "file.txt")) //menggabungkan elemen-elemen path menjadi satu path yang valid
}
