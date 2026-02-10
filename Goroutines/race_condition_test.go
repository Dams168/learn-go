package goroutines

/*
   Race Condition di paradigma Goroutine Go

   Race condition adalah kondisi di mana dua atau lebih goroutine mengakses data bersama secara bersamaan,
   dan hasil akhir tergantung pada urutan eksekusi goroutine tersebut.

   Race condition dapat menyebabkan perilaku yang tidak terduga dan sulit dideteksi dalam program konkuren.

*/

import (
	"fmt"
	"testing"
	"time"
)

func TestRaceCondition(t *testing.T) {
	x := 0

	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				x = x + 1
				// tanpa mekanisme sinkronisasi, ini dapat menyebabkan race condition
				// karena beberapa goroutine dapat membaca dan menulis nilai x secara bersamaan
				// yang mengakibatkan nilai akhir x tidak sesuai dengan yang diharapkan
				// yaitu 1000 * 100 = 100000
				// namun karena race condition, nilai akhir x bisa kurang dari itu
				// atau bahkan lebih besar dalam beberapa kasus ekstrem
			}
		}()
	}

	time.Sleep(5 * time.Second)

	fmt.Println("Final x = ", x)

}
