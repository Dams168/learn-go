package main

import (
	"encoding/csv"
	"os"
)

func main() {
	// Membuat writer CSV yang menulis ke output standar (os.Stdout)
	// os.Stdout adalah representasi dari output standar (biasanya terminal atau console)
	// csv.NewWriter membutuhkan io.Writer sebagai inputnya
	// kita bisa menggunakan os.Stdout karena os.Stdout mengimplementasikan interface io.Writer
	writer := csv.NewWriter(os.Stdout)

	// Menulis beberapa record (baris) ke output CSV
	// _ adalah cara untuk mengabaikan nilai kembalian dari fungsi Write
	_ = writer.Write([]string{"Nama", "Umur", "Kota"})
	_ = writer.Write([]string{"Andi", "30", "Jakarta"})
	_ = writer.Write([]string{"Budi", "25", "Bandung"})
	_ = writer.Write([]string{"Citra", "28", "Surabaya"})

	// Memastikan semua data tertulis ke output
	writer.Flush()
}
