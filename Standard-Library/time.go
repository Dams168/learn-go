package main

/*
   Package time di Go
   Digunakan untuk mengelola waktu dan tanggal, termasuk pengukuran durasi, format waktu, dan zona waktu.
   package ini sering digunakan dalam aplikasi yang memerlukan penanganan waktu, seperti penjadwalan tugas, pencatatan waktu, dan pengukuran interval waktu.

*/

import (
	"fmt"
	"time"
)

func main() {
	var localTime time.Time = time.Now()
	fmt.Println(localTime)

	utcTime := time.Date(2020, 10, 10, 10, 10, 10, 0, time.UTC)
	fmt.Println(utcTime)

	layout := "2006-01-02 15:04:05"
	valueTime := "2023-10-10 10:10:10"
	parsedTime, err := time.Parse(layout, valueTime)
	if err != nil {
		fmt.Println("Error parsing time:", err)
	} else {
		fmt.Println(parsedTime)
	}

	fmt.Println(parsedTime.Year())
	fmt.Println(parsedTime.Month())
	fmt.Println(parsedTime.Day())
	fmt.Println(parsedTime.Second())
	fmt.Println(parsedTime.Hour())

}
