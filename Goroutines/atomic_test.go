package goroutines

/*
sync/atomic di Go
sync/atomic adalah paket di Go yang menyediakan fungsi-fungsi untuk melakukan operasi atomik pada variabel-variabel yang dapat diakses oleh beberapa goroutine secara bersamaan.
Operasi atomik adalah operasi yang tidak dapat diinterupsi oleh goroutine lain, sehingga memastikan konsistensi data saat diakses secara konkuren.

Paket sync/atomic menyediakan fungsi-fungsi seperti AddInt64, LoadInt64, StoreInt64, CompareAndSwapInt64, dan lain-lain untuk melakukan operasi atomik pada variabel-variabel bertipe int64, uint64, dan pointer.
Fungsi-fungsi ini memungkinkan kita untuk melakukan operasi seperti penambahan, pembacaan, penulisan, dan perbandingan nilai secara atomik tanpa perlu menggunakan mutex atau mekanisme sinkronisasi lainnya.
Paket sync/atomic sangat berguna untuk menghindari race condition dan memastikan integritas data saat bekerja dengan goroutine yang mengakses variabel bersama.

*/

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestAtomic(t *testing.T) {
	var x int64 = 0
	group := sync.WaitGroup{}

	for i := 1; i <= 1000; i++ {
		go func() {
			/**
			Jika terjadi error : panic: sync: WaitGroup is reused before previous Wait has returned
			Itu artinya, goroutine belum selesai menjalankan kode group.Add(1), naun goroutine unit test
			sudah melakukan group.Wait(), group tidak boleh di add ketika sudah di Wait(), hal ini biasanya
			terjadi jika resource hardware kurang cepat ketika menjalankan goroutine diawal
			Jika hal ini terjadi, silahkan pindahkan kode group.Add(1), ke baris 15 sebelum memanggil go func()
			*/
			group.Add(1)
			for j := 1; j <= 100; j++ {
				atomic.AddInt64(&x, 1)
			}
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Counter = ", x)
}
