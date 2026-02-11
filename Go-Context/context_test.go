package gocontext

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
}

func TestContextWithValue(t *testing.T) {
	// Parent Context
	contextA := context.Background()

	//Child Context
	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	contextF := context.WithValue(contextC, "f", "F")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)

	// Dia akan mengakses nilai dari parent context, dalam artian child context bisa mengakses nilai dari parent context
	// Namun parent context tidak bisa mengakses nilai dari child context
	// Yang berarti hirarki keatas saja
	fmt.Println(contextF.Value("f"))
	fmt.Println(contextF.Value("b"))
	fmt.Println(contextF.Value("c"))
}

// Fungsi CreateCounter membuat sebuah channel yang mengirimkan angka secara berurutan setiap detik
// Fungsi ini menerima context sebagai parameter untuk mengontrol penghentian pengiriman angka
// Jika context dibatalkan, goroutine akan berhenti mengirim angka dan menutup channel
// Fungsi ini mengembalikan channel yang dapat digunakan untuk menerima angka dari goroutine
// Jika tidak menggunakan context, goroutine akan terus berjalan tanpa henti yang menyebabkan goroutine leak
func CreateCounter(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1
		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
				time.Sleep(1 * time.Second)
			}
		}
	}()

	return destination
}

/*
	Context di golang digunakan untuk mengelola batas waktu, pembatalan, dan nilai-nilai yang terkait dengan permintaan atau operasi tertentu.
	Disini ada beberapa fungsi yang menunjukkan penggunaan context dengan berbagai cara:

	1. TestContextWithCancel: Fungsi ini membuat context dengan kemampuan untuk dibatalkan menggunakan context.WithCancel.
	   Setelah menerima angka hingga 10, context dibatalkan yang menghentikan goroutine yang mengirim angka.
	   Context ini digunakan ketika kita ingin menghentikan operasi secara eksplisit, yang berarti goroutine tidak akan terus berjalan tanpa henti.
	2. TestContextWithTimeout: Fungsi ini membuat context dengan batas waktu menggunakan context.WithTimeout.
	   Context ini secara otomatis dibatalkan setelah 5 detik, yang menghentikan goroutine jika belum selesai.
	   Context ini digunakan ketika kita ingin memastikan bahwa operasi tidak berjalan lebih lama dari waktu yang ditentukan, yang berarti goroutine akan dihentikan setelah waktu habis.
	3. TestContextWithDeadline: Fungsi ini mirip dengan TestContextWithTimeout, tetapi menggunakan context.WithDeadline untuk menetapkan waktu kedaluwarsa tertentu.
	   Setelah waktu tersebut tercapai, context dibatalkan dan goroutine dihentikan.

*/

func TestContextWithCancel(t *testing.T) {
	fmt.Println("Total Goroutine : ", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	destination := CreateCounter(ctx)

	for n := range destination {
		fmt.Println("Counter :", n)
		if n == 10 {
			break
		}
	}

	cancel()
	time.Sleep(2 * time.Second)

	fmt.Println("Total Goroutine : ", runtime.NumGoroutine())
}

func TestContextWithTimeout(t *testing.T) {

	fmt.Println("Total Goroutine : ", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()

	destination := CreateCounter(ctx)

	for n := range destination {
		fmt.Println("Counter :", n)
		if n == 10 {
			break
		}
	}

	fmt.Println("Total Goroutine : ", runtime.NumGoroutine())
}

func TestContextWithDeadline(t *testing.T) {

	fmt.Println("Total Goroutine : ", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(5*time.Second))
	defer cancel()

	destination := CreateCounter(ctx)

	for n := range destination {
		fmt.Println("Counter :", n)
		if n == 10 {
			break
		}
	}

	fmt.Println("Total Goroutine : ", runtime.NumGoroutine())
}
