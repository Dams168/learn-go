package main

/*
   Flag package di Go
   Digunakan untuk parsing argumen command line,
   sehingga program dapat menerima input dari user saat dijalankan.

*/

import (
	"flag"
	"fmt"
)

func main() {
	// Mendefinisikan flag/argumen command line, kenapa menggunakan pointer?
	// Karena flag.String dan flag.Int mengembalikan pointer ke nilai yang di-parse
	var username *string = flag.String("username", "root", "database username")
	var password *string = flag.String("password", "root", "database password")
	var host *string = flag.String("host", "localhost", "database host")
	var port *int = flag.Int("port", 3306, "database port")

	// Mem-parsing argumen command line
	// Setelah flag.Parse() dipanggil, variabel-variabel di atas akan berisi
	// nilai yang diambil dari argumen command line jika ada, atau nilai default jika tidak ada
	flag.Parse()

	fmt.Println("username :", *username)
	fmt.Println("password :", *password)
	fmt.Println("host :", *host)
	fmt.Println("port :", *port)
}
