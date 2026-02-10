package goroutines

/*
sync.Cond di Go
sync.Cond adalah struktur data di Go yang digunakan untuk mengoordinasikan goroutine yang menunggu kondisi tertentu sebelum melanjutkan eksekusi.
sync.Cond menyediakan mekanisme untuk menunggu (Wait) dan memberi sinyal (Signal atau Broadcast) kepada goroutine yang menunggu.

Fungsi utama dari sync.Cond adalah:
- Wait(): Menyebabkan goroutine yang
    memanggilnya untuk menunggu hingga diberi sinyal.
- Signal(): Memberi sinyal kepada satu goroutine yang menunggu untuk melanjutkan eksekusi.
- Broadcast(): Memberi sinyal kepada semua goroutine yang menunggu untuk melanjutkan eksekusi.

Ketika menggunakan sync.Cond, biasanya diperlukan mutex (sync.Mutex atau sync.RWMutex) untuk melindungi akses ke kondisi bersama yang sedang dipantau.
*/

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var locker = sync.Mutex{}
var cond = sync.NewCond(&locker)
var group = sync.WaitGroup{}

func WaitCondition(value int) {
	defer group.Done()
	group.Add(1)

	cond.L.Lock()
	cond.Wait()
	fmt.Println("Done", value)
	cond.L.Unlock()
}

func TestCond(t *testing.T) {
	for i := 0; i < 10; i++ {
		go WaitCondition(i)
	}

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			cond.Signal()
		}
	}()

	//go func() {
	//	time.Sleep(1 * time.Second)
	//	cond.Broadcast()
	//}()

	group.Wait()
}
