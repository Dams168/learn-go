package goroutines

/*
   Test Ini digunakan Untuk Mencegah Race Condition dengan Mutex dan RWMutex

*/

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/*
   Mutex di paradigma Goroutine Go
   Mutex (mutual exclusion) adalah mekanisme sinkronisasi yang digunakan untuk melindungi akses ke sumber daya bersama
   agar hanya satu goroutine yang dapat mengaksesnya pada satu waktu.
   Mutex digunakan untuk mencegah race condition dengan cara mengunci sumber daya sebelum diakses dan membukanya kembali setelah selesai digunakan.

   Mutex di Go disediakan oleh paket "sync" dan memiliki dua metode utama: Lock() dan Unlock().
   - Lock(): Mengunci mutex. Jika mutex sudah terkunci oleh goroutine lain, goroutine yang memanggil Lock() akan menunggu hingga mutex tersedia.
   - Unlock(): Membuka kunci mutex, memungkinkan goroutine lain untuk mengakses sumber daya bersama.

*/

func TestMutex(t *testing.T) {
	x := 0
	var mutex sync.Mutex

	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				mutex.Lock()
				x = x + 1
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)

	fmt.Println("Final x = ", x)
}

/*
   RWMutex di paradigma Goroutine Go
   RWMutex (Read-Write Mutex) adalah varian dari Mutex yang memungkinkan beberapa goroutine untuk membaca sumber daya bersama secara bersamaan,
   tetapi hanya satu goroutine yang dapat menulis ke sumber daya tersebut pada satu waktu.
   RWMutex berguna ketika ada lebih banyak operasi baca daripada operasi tulis, sehingga meningkatkan kinerja dengan memungkinkan akses baca paralel.

   RWMutex memiliki dua jenis kunci:
   - RLock(): Mengunci RWMutex untuk operasi baca. Beberapa goroutine dapat memegang kunci baca secara bersamaan.
   - Lock(): Mengunci RWMutex untuk operasi tulis. Hanya satu goroutine yang dapat memegang kunci tulis, dan tidak ada goroutine lain yang dapat memegang kunci baca atau tulis saat kunci tulis dipegang.
   - RUnlock(): Membuka kunci baca RWMutex.
   - Unlock(): Membuka kunci tulis RWMutex.

   RWMutex digunakan untuk meningkatkan efisiensi akses ke sumber daya bersama ketika operasi baca lebih sering terjadi dibandingkan operasi tulis.
   RWMutex memungkinkan banyak goroutine untuk membaca data secara bersamaan tanpa saling menghalangi, selama tidak ada goroutine yang menulis data pada saat yang sama.

*/

type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (account *BankAccount) AddBalance(amount int) {
	account.RWMutex.Lock()
	account.Balance = account.Balance + amount
	account.RWMutex.Unlock()
}

func (account *BankAccount) GetBalance() int {
	account.RWMutex.Lock()
	balance := account.Balance
	account.RWMutex.Unlock()
	return balance
}

func TestRWMutex(t *testing.T) {
	account := BankAccount{}
	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				account.AddBalance(1)
			}
		}()
	}

	time.Sleep(5 * time.Second)

	fmt.Println("Total Balance = ", account.GetBalance())
}

type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

// Fungsi Lock ini adalah method dari struct UserBalance yang digunakan untuk mengunci mutex yang ada di dalam struct tersebut.
func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

// Fungsi Unlock ini adalah method dari struct UserBalance yang digunakan untuk membuka kunci mutex yang ada di dalam struct tersebut.
func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}
func (user *UserBalance) Change(amount int) {
	user.Balance = user.Balance + amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int) {
	// pertama-tama mengunci user1, kenapa? karena kita akan mengurangi saldo user1
	// setelah itu kita mengunci user2, karena kita akan menambahkan saldo user2
	// dengan mengunci kedua user tersebut, kita memastikan tidak ada goroutine lain yang bisa mengubah saldo mereka selama proses transfer berlangsung
	user1.Lock()
	fmt.Println("Lock User 1 : ", user1.Name)
	user1.Change(-amount)

	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("Lock User 2 : ", user2.Name)
	user2.Change(amount)

	time.Sleep(3 * time.Second)

	user1.Unlock()
	user2.Unlock()
}

func TestDeadLock(t *testing.T) {
	user1 := UserBalance{Name: "Dam", Balance: 1000000}
	user2 := UserBalance{Name: "Dim", Balance: 1000000}

	// goroutine pertama mencoba mentransfer uang dari user1 ke user2
	// goroutine kedua mencoba mentransfer uang dari user2 ke user1
	// namun karena urutan penguncian yang berbeda, ini dapat menyebabkan deadlock
	// di mana kedua goroutine saling menunggu untuk membuka kunci yang dipegang oleh yang lain
	// sehingga kedua goroutine tidak pernah bisa melanjutkan eksekusi mereka
	// untuk menghindari deadlock, kita harus memastikan bahwa semua goroutine mengunci sumber daya dalam urutan yang konsisten
	// misalnya, selalu mengunci user dengan nama yang lebih kecil terlebih dahulu
	// dalam contoh ini, kita tidak menerapkan solusi tersebut, sehingga kemungkinan besar akan terjadi deadlock
	go Transfer(&user1, &user2, 100000)
	go Transfer(&user2, &user1, 200000)

	time.Sleep(10 * time.Second)

	fmt.Println("User 1 :", user1.Name, " Balance : ", user1.Balance)
	fmt.Println("User 2 :", user2.Name, " Balance : ", user2.Balance)
}
