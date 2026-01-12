package main

// interface adalah tipe data abstrak yang mendefinisikan kumpulan method tanpa implementasi
// interface digunakan untuk mendefinisikan perilaku yang harus dimiliki oleh tipe data tertentu
// sebuah tipe data dikatakan mengimplementasikan interface jika tipe data tersebut memiliki semua method yang didefinisikan dalam interface tersebut

type hitung interface {
	luas() int
	keliling() int
}

type persegiPanjang struct {
	panjang int
	lebar   int
}

// implementasi method dari interface hitung untuk tipe data persegiPanjang
func (p persegiPanjang) luas() int {
	return p.panjang * p.lebar
}

func (p persegiPanjang) keliling() int {
	return 2 * (p.panjang + p.lebar)
}

// contoh fungsi yang mengembalikan tipe data interface{}
// fungsi ini dapat mengembalikan nilai dari tipe data apapun
// dalam contoh ini, fungsi ups mengembalikan string
func ups() interface{} {
	return "ups"
}

func main() {
	// membuat instance dari persegiPanjang dan menggunakannya sebagai hitung
	var h hitung = persegiPanjang{panjang: 10, lebar: 5}

	println("Luas:", h.luas())
	println("Keliling:", h.keliling())
	println("Ups:", ups().(string)) // melakukan type assertion untuk mengubah interface{} menjadi string

}
