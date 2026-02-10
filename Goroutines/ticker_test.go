package goroutines

/*
Ticker di Go

Ticker adalah mekanisme yang digunakan untuk menjalankan kode secara berulang pada interval waktu tertentu.
Go menyediakan dua cara utama untuk bekerja dengan ticker, yaitu time.NewTicker dan time.Tick.

- time.NewTicker: Membuat ticker baru yang mengirimkan waktu saat ticker berdetak ke channel C pada interval yang ditentukan.
- time.Tick: Mengembalikan channel yang menerima waktu saat ticker berdetak pada interval yang ditentukan.

Ticker berguna untuk tugas-tugas yang memerlukan eksekusi berkala, seperti polling, pembaruan status, atau tugas terjadwal lainnya.
*/

import (
	"fmt"
	"testing"
	"time"
)

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)

	go func() {
		time.Sleep(5 * time.Second)
		ticker.Stop()
	}()

	for time := range ticker.C {
		fmt.Println(time)
	}
}

func TestTick(t *testing.T) {
	channel := time.Tick(1 * time.Second)

	for time := range channel {
		fmt.Println(time)
	}
}
