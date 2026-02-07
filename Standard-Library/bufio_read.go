package main

/*
package bufio di Go
Digunakan untuk melakukan buffered I/O (Input/Output) pada data.
Package ini menyediakan fungsi dan tipe data untuk membaca dan menulis data dengan efisiensi yang lebih tinggi melalui penggunaan buffer.
Buffered I/O membantu mengurangi jumlah operasi I/O langsung ke sumber data, yang dapat meningkatkan kinerja aplikasi, terutama saat bekerja dengan file atau jaringan.

*/

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	input := strings.NewReader("this is long string\nwith new line\n")
	reader := bufio.NewReader(input)

	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}

		fmt.Println(string(line))
	}
}
