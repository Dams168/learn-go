package main

import (
	"fmt"
	"path"
)

func main() {
	fmt.Println(path.Dir("hello/world.go"))              //mengembalikan direktori dari path yang diberikan
	fmt.Println(path.Base("hello/world.go"))             //mengembalikan nama file dari path yang diberikan
	fmt.Println(path.Ext("hello/world.go"))              //mengembalikan ekstensi file dari path yang diberikan
	fmt.Println(path.Join("hello", "world", "file.txt")) //menggabungkan elemen-elemen path menjadi satu path yang valid
}
