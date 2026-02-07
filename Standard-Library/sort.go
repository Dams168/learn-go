package main

/*
   Package Sort Go
   Package sort ini berfungsi untuk mengurutkan data dalam berbagai struktur seperti slice dan user-defined types.
   Pustaka ini menyediakan berbagai fungsi dan metode untuk mengurutkan data secara efisien, baik dalam urutan menaik maupun menurun.
   Selain itu, package ini juga memungkinkan pengguna untuk mengimplementasikan kriteria pengurutan khusus dengan mendefinisikan interface sort.Interface.

    Kontrak sort.Interface terdiri dari tiga metode utama:
    - Len() int: Mengembalikan jumlah elemen dalam koleksi.
    - Swap(i, j int): Menukar elemen pada indeks i dan j.
    - Less(i, j int) bool: Menentukan apakah elemen pada indeks i harus diurutkan sebelum elemen pada indeks j.

\
*/

import (
	"fmt"
	"sort"
)

type users struct {
	name string
	age  int
}

// Implementasi interface sort.Interface untuk userSlice
// Dengan mengimplementasikan metode Len, Swap, dan Less
type userSlice []users

// Kontrak sort.Interface
func (u userSlice) Len() int {
	return len(u)
}

func (u userSlice) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

func (u userSlice) Less(i, j int) bool {
	return u[i].age < u[j].age
}

func main() {
	datas := userSlice{
		{"Dam", 22},
		{"Andi", 20},
		{"Budi", 25},
		{"Cici", 19},
	}

	sort.Sort(userSlice(datas))

	fmt.Println(datas)
}
