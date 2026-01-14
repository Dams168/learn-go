package main

import "fmt"

func endApplication() {
	fmt.Println("End Application")
}

func runApplication() {
	defer endApplication() // defer akan menunda pemanggilan function sampai function yang memanggilnya selesai dieksekusi
	fmt.Println("Run Application")
}

func main() {
	runApplication()
}
