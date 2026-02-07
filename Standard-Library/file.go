package main

/*
file manipulation di go
Digunakan untuk melakukan operasi pada file seperti membuat file baru, menulis ke file,
menambahkan konten ke file yang sudah ada, dan membaca konten dari file.

Package ini menyediakan fungsi untuk membuka, menutup, membaca, dan menulis file dengan berbagai mode akses.
Operasi file adalah bagian penting dalam pengembangan aplikasi yang memerlukan penyimpanan data secara persisten.

*/

import (
	"bufio"
	"io"
	"os"
)

func createNewFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(message)
	return nil
}

func addToFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(message)
	return nil
}

func readFile(name string) (string, error) {
	file, err := os.OpenFile(name, os.O_RDONLY, 0666)
	if err != nil {
		return "", err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var message string
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		message += string(line) + "\n"
	}
	return message, nil
}

func main() {
	// createNewFile("sample.log", "this is sample log")

	// result, _ := readFile("sample.log")
	// fmt.Println(result)

	addToFile("sample.log", "\nthis is add message")
}
