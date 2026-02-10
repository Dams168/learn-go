package goroutines

/*
channel di Go
Channel adalah mekanisme komunikasi antar goroutine di Go. Channel memungkinkan goroutine untuk saling bertukar data dengan cara yang aman dan terkoordinasi.
Channel dibuat menggunakan fungsi make dengan sintaks make(chan T), di mana T adalah tipe data yang akan dikirim melalui channel.

Untuk mengirim data ke channel, kita menggunakan operator <- dengan sintaks channel <- value.
Untuk menerima data dari channel, kita juga menggunakan operator <- dengan sintaks value := <-channel.

Channel dapat bersifat unbuffered (tanpa buffer) atau buffered (dengan buffer).
Channel unbuffered mengharuskan pengirim dan penerima untuk siap secara bersamaan agar data dapat dikirim.
Channel buffered memungkinkan pengirim untuk mengirim sejumlah data tertentu tanpa menunggu penerima, hingga buffer penuh.



*/

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Dam in the channel"
		fmt.Println("Finish Send to Channel")
	}()

	data := <-channel
	fmt.Println(data)

	time.Sleep(4 * time.Second)
}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Dam Join Channel"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)
}

func ChannelIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Dam Join Channel"
}

func ChannelOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)
	go ChannelIn(channel)
	go ChannelOut(channel)
	time.Sleep(3 * time.Second)
}

/*
Seperti yang dijelaskan sebelumnya, bahwa secara default channel itu hanya bisa menerima 1 data
Artinya jika kita menambah data ke-2, maka kita akan diminta menunggu sampai data ke-1 ada yang mengambil
Kadang-kadang ada kasus dimana pengirim lebih cepat dibanding penerima, dalam hal ini jika kita menggunakan channel, maka otomatis pengirim akan ikut lambat juga
Untuknya ada Buffered Channel, yaitu buffer yang bisa digunakan untuk menampung data antrian di Channel

Kita bebas memasukkan berapa jumlah kapasitas antrian di dalam buffer
Jika kita set misal 5, artinya kita bisa menerima 5 data di buffer.
Jika kita mengirim data ke 6, maka kita diminta untuk menunggu sampai buffer ada yang kosong
Ini cocok sekali ketika memang goroutine yang menerima data lebih lambat dari yang mengirim data
*/

func TestBufferChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	go func() {
		channel <- "Dam 1"
		channel <- "Dam 2"
		channel <- "Dam 3"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(3 * time.Second)
	fmt.Println("Finish")
}

/*
Kadang-kadang ada kasus sebuah channel dikirim data secara terus menerus oleh pengirim
Dan kadang tidak jelas kapan channel tersebut akan berhenti menerima data
Salah satu yang bisa kita lakukan adalah dengan menggunakan perulangan range ketika menerima data dari channel
Ketika sebuah channel di close(), maka secara otomatis perulangan tersebut akan berhenti
Ini lebih sederhana dari pada kita melakukan pengecekan channel secara manual
*/
func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- fmt.Sprintf("Message %d", i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println(data)
	}

	fmt.Println("Finish")
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Received from channel 1:", data)
			counter++
		case data := <-channel2:
			fmt.Println("Received from channel 2:", data)
			counter++
		}
		if counter == 2 {
			break
		}
	}
}

func TestDefaultSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Received from channel 1:", data)
			counter++
		case data := <-channel2:
			fmt.Println("Received from channel 2:", data)
			counter++
		default:
			fmt.Println("Waiting for data...")
		}
		if counter == 2 {
			break
		}
	}
}
