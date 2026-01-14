package main

import "fmt"

/*
   Method dengan Receiver Pointer
   Method dengan receiver pointer adalah method yang memiliki receiver bertipe pointer.
   Dengan menggunakan receiver pointer, kita dapat mengubah nilai dari struct yang memanggil method tersebut.
   karena receiver pointer menunjuk langsung ke alamat memori dari struct tersebut.
   Direkomendasikan menggunakan pointer receiver ketika method perlu mengubah nilai dari struct
   atau ketika struct tersebut berukuran besar untuk menghindari overhead penyalinan nilai struct.

*/

type Man struct {
	Name string
}

// Married adalah method dengan receiver pointer, sehingga dapat mengubah nilai dari struct Man yang memanggilnya
// Jika tidak menggunakan pointer receiver, perubahan pada Name tidak akan mempengaruhi struct
func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	josh := Man{"Josh"}
	josh.Married()
	fmt.Println(josh.Name)
}
