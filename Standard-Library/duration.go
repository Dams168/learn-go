package main

/*
   Package duration di Go
   Digunakan untuk merepresentasikan dan mengelola durasi waktu.
   Tipe data utama dalam package ini adalah time.Duration, yang merepresentasikan selisih waktu dalam nanodetik.
   Package ini menyediakan berbagai fungsi dan metode untuk melakukan operasi pada durasi, seperti penambahan, pengurangan, dan konversi antara berbagai satuan waktu (detik, menit, jam, dll).
*/

import (
	"fmt"
	"time"
)

func main() {
	var dur1 time.Duration = 100 * time.Second
	var dur2 time.Duration = 2 * time.Hour

	totalDuration := dur1 + dur2

	fmt.Println("Total Duration:", totalDuration)
}
