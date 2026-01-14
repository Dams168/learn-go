package main

import "fmt"

/*
   Fungsi dengan Parameter Pointer
   Fungsi dengan parameter pointer adalah fungsi yang menerima argumen berupa pointer
   yang menunjuk ke lokasi memori dari sebuah variable.
   Dengan menggunakan pointer sebagai parameter, kita dapat mengubah nilai dari variable
   yang berada di luar fungsi tersebut.
   karena fungsi menerima alamat memori dari variable tersebut,
   sehingga perubahan yang dilakukan pada nilai melalui pointer akan mempengaruhi nilai asli dari variable tersebut.


*/

type address struct {
	city, province, country string
}

// changeAddress mengubah nilai city dari address melalui pointer
func changeAddress(addr *address) {
	addr.city = "Cianjur"
}

func main() {
	// membuat variable address1 bertipe address
	address1 := address{}
	changeAddress(&address1) // menggunakan & untuk mendapatkan alamat memori dari address1
	fmt.Println(address1)

}
