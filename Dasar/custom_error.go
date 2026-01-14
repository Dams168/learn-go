package main

/*
Custom Error Handling in Go
Di Go, kita dapat membuat custom error dengan cara membuat tipe data baru yang mengimplementasikan interface error.
Interface error memiliki satu metode yaitu Error() yang mengembalikan string yang berisi pesan kesalahan.

alurnya adalah sebagai berikut:
1. Buat tipe data baru untuk custom error.
2. Implementasikan metode Error() untuk tipe data tersebut.
3. Gunakan custom error tersebut di dalam fungsi atau metode sesuai kebutuhan.
4. Saat menangani error, lakukan type assertion untuk memeriksa tipe custom error dan tangani sesuai kebutuhan.


*/

type validationError struct {
	message string
}

type notFoundError struct {
	message string
}

// implementasi metode Error() untuk custom error, sehingga tipe data ini mengimplementasikan interface error
// harus menggunakan pointer receiver agar perubahan pada struct dapat terlihat saat diakses diluar
// jika menggunakan value receiver, perubahan pada struct tidak akan terlihat diluar karena yang dioper adalah salinan dari struct tersebut
func (e *validationError) Error() string {
	return e.message
}

// implementasi metode Error() untuk custom error, sehingga tipe data ini mengimplementasikan interface error
func (e *notFoundError) Error() string {
	return e.message
}

// fungsi yang menggunakan custom error
func saveData(id string, data any) error {

	if id == "" {
		// Harus mengembalikan pointer ke struct validationError, karena metode Error() diimplementasikan dengan pointer receiver
		// jika tidak menggunakan pointer, maka akan terjadi error saat melakukan type assertion
		return &validationError{message: "ID cannot be empty"}
	}

	if id != "dam" {
		return &notFoundError{message: "Data not found"}
	}
	// jika tidak ada error, kembalikan nil
	return nil
}

func main() {
	err := saveData("", nil)

	if err != nil {
		// switch e := err.(type) {
		// case *validationError:
		// 	println("Validation Error:", e.Error())
		// case *notFoundError:
		// 	println("Not Found Error:", e.Error())
		// default:
		// 	println("Unknown Error:", e.Error())
		// }

		//menggunakan if else untuk type assertion
		// Pengecekan error ini berguna ketika kita ingin menangani error berdasarkan tipe custom error yang dibuat
		// dimana pemanggil fungsi dapat mengetahui jenis error yang terjadi dan menanganinya sesuai kebutuhan
		if validationError, ok := err.(*validationError); ok {
			// jika err adalah tipe *validationError
			println("Validation Error:", validationError.Error())
		} else if notFoundError, ok := err.(*notFoundError); ok {
			println("Not Found Error:", notFoundError.Error())
		} else {
			println("Unknown Error:", err.Error())
		}
	} else {
		println("Data saved successfully")
	}
}
