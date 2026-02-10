package goroutines

/*
Timer di Go
Timer adalah mekanisme yang digunakan untuk menunda eksekusi suatu fungsi atau blok kode hingga waktu tertentu telah berlalu.
Go menyediakan beberapa cara untuk bekerja dengan timer, termasuk time.NewTimer, time.After, dan time.AfterFunc.
- time.NewTimer: Membuat timer baru yang akan mengirimkan waktu saat timer selesai ke channel C setelah durasi tertentu.
- time.After: Mengembalikan channel yang akan menerima waktu saat timer selesai setelah durasi tertentu.
- time.AfterFunc: Menjadwalkan fungsi untuk dijalankan setelah durasi tertentu.
Timer sangat berguna untuk mengatur waktu tunggu, menunda eksekusi, atau membuat batas waktu (timeout) dalam program Go.
*/

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)
	fmt.Println(time.Now())

	time := <-timer.C
	fmt.Println(time)
}

func TestAfter(t *testing.T) {
	channel := time.After(5 * time.Second)
	fmt.Println(time.Now())

	time := <-channel
	fmt.Println(time)
}

func TestAfterFunc(t *testing.T) {
	group := sync.WaitGroup{}
	group.Add(1)

	time.AfterFunc(5*time.Second, func() {
		fmt.Println(time.Now())
		group.Done()
	})
	fmt.Println(time.Now())

	group.Wait()
}
