package goroutines

/*
sync.Pool di Go
sync.Pool adalah struktur data di Go yang digunakan untuk mengelola kumpulan objek yang dapat digunakan kembali (reusable objects).

Implementasi Pool sudah aman dari Race Condition

Objek-objek ini dapat diambil dari pool ketika dibutuhkan dan dikembalikan ke pool setelah selesai digunakan.
Tujuan utama dari sync.Pool adalah untuk mengurangi overhead alokasi memori dengan cara mendaur ulang objek-objek yang sudah ada,
sehingga meningkatkan kinerja aplikasi, terutama dalam situasi di mana objek dibuat dan dihancurkan secara berulang-ulang.

sync.Pool memiliki dua metode utama:
- Get(): Mengambil objek dari pool. Jika pool kosong, metode ini akan memanggil fungsi New (jika disediakan) untuk membuat objek baru.
- Put(x any): Mengembalikan objek x ke pool, sehingga dapat digunakan kembali di masa depan.

Keunggulan sync.Pool adalah kemampuannya untuk secara otomatis mengelola siklus hidup objek, termasuk pembersihan objek yang tidak lagi digunakan oleh garbage collector,
yang membantu mengoptimalkan penggunaan memori dalam aplikasi Go.

Kekurangan sync.Pool adalah bahwa objek yang disimpan di dalamnya tidak dijamin untuk tetap ada,
karena garbage collector dapat menghapus objek-objek tersebut jika diperlukan untuk mengelola memori.
Oleh karena itu, sync.Pool lebih cocok untuk objek-objek yang bersifat sementara dan dapat dibuat ulang dengan mudah.

*/

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	pool := sync.Pool{
		// Fungsi untuk membuat objek baru jika pool kosong
		New: func() any {
			return "New"
		},
	}
	// Menambahkan data ke pool
	pool.Put("Dam")
	pool.Put("Dim")
	pool.Put("Dum")

	for i := 0; i < 10; i++ {
		go func() {
			// Mengambil data dari pool
			data := pool.Get()
			fmt.Println(data)

			// Sleep untuk mensimulasikan proses penggunaan data
			time.Sleep(1 * time.Second)
			// Mengembalikan data ke pool
			pool.Put(data)
		}()
	}

	time.Sleep(11 * time.Second)
	fmt.Println("Selesai")
}
