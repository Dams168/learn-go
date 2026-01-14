package main

/*
Error Handling di Go

Di Go, error handling dilakukan dengan mengembalikan nilai error dari sebuah fungsi.
Jika terjadi kesalahan, fungsi akan mengembalikan nilai error yang dapat diperiksa oleh pemanggil fungsi.

Cara umum untuk membuat error di Go adalah dengan menggunakan package "errors" yang menyediakan fungsi New untuk membuat error baru.
ini memungkinkan kita untuk membuat pesan error yang spesifik sesuai kebutuhan aplikasi kita.

Interface error di Go adalah tipe data built-in yang digunakan untuk merepresentasikan kesalahan.
interface error memiliki satu metode yaitu Error() yang mengembalikan string yang berisi pesan kesalahan.

Contoh penggunaan error handling di Go ditunjukkan pada fungsi pembagian di bawah ini.
Fungsi ini menerima dua parameter yaitu nilai dan pembagi, dan mengembalikan hasil pembagian serta nilai error.
Jika pembagi adalah nol, fungsi akan mengembalikan error baru dengan pesan "Pembagian dengan nol".
Jika tidak, fungsi akan mengembalikan hasil pembagian dan nil sebagai nilai error.

*/

import (
	"errors"
	"fmt"
)

func pembagian(nilai int, pembagi int) (hasil int, err error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagian dengan nol")
	} else {
		return nilai / pembagi, nil
	}
}

func main() {
	hasil, err := pembagian(10, 2)
	if err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Println("Hasil:", hasil)
	}

}
