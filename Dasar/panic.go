package main

import "fmt"

/*
   Panic, Recover, Defer
   - Panic digunakan untuk menghentikan program secara paksa ketika terjadi error yang tidak terduga
   - Recover digunakan untuk menangkap panic dan mencegah program berhenti secara paksa
   - Defer digunakan untuk menunda eksekusi function sampai function yang memanggilnya selesai dieksekusi

   fungsi ini mirip dengan try catch pada bahasa pemrograman lain, dimana panic berfungsi sebagai throw dan recover berfungsi sebagai catch
   defer berfungsi untuk memastikan bahwa recover akan selalu dieksekusi meskipun terjadi panic

*/

func endApp() {
	fmt.Println("End App")

	msg := recover() // recover digunakan untuk menangkap panic dan mencegah program berhenti secara paksa
	if msg != nil {
		// jika ada panic, maka msg tidak nil
		fmt.Println("Error terjadi:", msg)
	}
}

func runApp(error bool) {
	// defer tetap menjalankan function endApp walaupun terjadi panic
	defer endApp()
	if error {
		// panic digunakan untuk menghentikan program secara paksa ketika terjadi error, biasanya digunakan saat terjadi error yang tidak terduga
		// ketika panic terjadi, program akan langsung berhenti dan mengeksekusi semua defer yang ada sebelum benar-benar berhenti
		// ini berguna untuk membersihkan resource atau melakukan logging sebelum program berhenti, sehingga tidak ada resource yang bocor atau data yang hilang
		panic("Yah error Nich...")
	}

	//program tidak akan mencapai titik ini jika panic terjadi
	// contoh recover yang salah
	// msg := recover()
	// if msg != nil {
	// 	fmt.Println("Error terjadi:", msg)
	// }
}

func main() {
	runApp(true)
	//jika tidak ada recover di endApp, program akan berhenti dan tidak mengeksekusi baris di bawah ini
	//jadi pastikan menggunakan recover di defer function untuk menangani panic agar program tetap berjalan
	fmt.Println("App Running")
}
