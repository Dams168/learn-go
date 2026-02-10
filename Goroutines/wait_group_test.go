package goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsynchronous(group *sync.WaitGroup) {
	// menandai bahwa goroutine ini telah selesai ketika fungsi ini keluar
	// jika tidak menggunakan Done, maka program akan menunggu selamanya pada group.Wait() dan terjadi deadlock
	defer group.Done()

	// Simulasi pekerjaan yang memakan waktu
	// setiap kali goroutine dijalankan, kita menambahkan 1 ke WaitGroup
	group.Add(1)

	fmt.Println("Async Task")

	time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	// membuat instance WaitGroup
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsynchronous(group)
	}

	// menunggu semua goroutine selesai, jika tidak ada yang memanggil Done, maka program akan menunggu selamanya
	// pada titik ini, program akan menunggu hingga semua goroutine yang telah menambahkan dirinya ke WaitGroup selesai
	group.Wait()

	fmt.Println("All Asynchronous Done")
}
