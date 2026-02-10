package goroutines

/*
sync.Once di Go
sync.Once adalah struktur data di Go yang digunakan untuk memastikan bahwa sebuah fungsi hanya dijalankan satu kali,
terlepas dari berapa kali fungsi tersebut dipanggil dari berbagai goroutine.

sync.Once memiliki metode Do(f func()), yang menerima sebuah fungsi sebagai parameter.
Metode Do akan menjalankan fungsi tersebut hanya pada pemanggilan pertama, dan mengabaikan pemanggilan berikutnya.
Ini sangat berguna untuk inisialisasi yang hanya perlu dilakukan sekali, seperti pengaturan konfigurasi atau pembuatan koneksi database.


*/

import (
	"fmt"
	"sync"
	"testing"
)

var counter = 0

func OnlyOnce() {
	counter++
}

func TestOnlyOnce(t *testing.T) {
	once := sync.Once{}
	group := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go func() {
			group.Add(1)
			once.Do(OnlyOnce)
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Counter ", counter)
}
