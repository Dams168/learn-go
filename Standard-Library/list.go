package main

/*
   Penggunaan Package list pada GoLang
   Package list menyediakan implementasi struktur data daftar ganda (doubly linked list).
   Pada double linked list, setiap elemen (node) memiliki referensi ke elemen sebelumnya dan berikutnya, sehingga memungkinkan traversal dua arah.
*/

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()
	data.PushBack("John")
	data.PushBack("Doe")
	data.PushBack("Dcaprio")
	data.PushFront("Leonardo")

	fmt.Println("Length of list:", data.Len())

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(string))
	}
}
