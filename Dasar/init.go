package main

import (
	"belajar-go-dasar/db"
	_ "belajar-go-dasar/internal" // import untuk menjalankan init function, walaupun tidak digunakan secara langsung, diperlukan untuk package internal
	"fmt"
)

func main() {
	fmt.Println(db.GetDatabase())
}
