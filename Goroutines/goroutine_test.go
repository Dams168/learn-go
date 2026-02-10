package goroutines

/*
Goroutines di Go
Goroutine adalah fungsi atau metode yang berjalan secara konkuren dengan fungsi atau metode lainnya.
Goroutine sangat ringan dan dikelola oleh runtime Go, memungkinkan ribuan goroutine berjalan secara bersamaan tanpa overhead yang signifikan.

Untuk membuat goroutine, kita cukup menambahkan kata kunci "go" sebelum pemanggilan fungsi atau metode.
Ini akan menjalankan fungsi tersebut secara asinkron, memungkinkan program untuk melanjutkan eksekusi tanpa menunggu fungsi tersebut selesai.
*/

import (
	"fmt"
	"testing"
	"time"
)

func HelloWorld() {
	fmt.Println("Hello World!")
}

func TestGoroutines(t *testing.T) {
	go HelloWorld()
	fmt.Println("ups!")

	time.Sleep(1 * time.Second)
}

func DisplayNumber(number int) {
	fmt.Println("Number : ", number)
}

func TestManyGoroutines(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(5 * time.Second)
}
