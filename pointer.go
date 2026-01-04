package main

import "fmt"

// pointer = alamat memori dari sebuah variable
// penting untuk menghemat penggunaan memory,

// & digunakan untuk mendapatkan alamat memori dari sebuah variable
// * digunakan untuk mengakses value dari sebuah pointer atau mendeklarasikan sebuah variable sebagai pointer
func pointerExample() {
	var numberA int = 4
	var numberB *int = &numberA

	fmt.Println("numberA (value)   :", numberA)  // 4
	fmt.Println("numberA (address) :", &numberA) // 0xc20800a220

	fmt.Println("numberB (value)   :", *numberB) // 4
	fmt.Println("numberB (address) :", numberB)  // 0xc20800a220

	fmt.Println("")

	numberA = 5

	// perubahan value pada numberA akan berpengaruh pada numberB
	// karena numberB adalah pointer yang mereferensikan alamat memori dari numberA
	// sehingga ketika value pada alamat memori tersebut berubah, maka value yang diakses melalui pointer juga berubah
	fmt.Println("numberA (value)   :", numberA)
	fmt.Println("numberA (address) :", &numberA)
	fmt.Println("numberB (value)   :", *numberB)
	fmt.Println("numberB (address) :", numberB)

}

func withoutPointer() {

	x := 10
	x2 := x

	x2 = 20

	// nilai x tidak berubah karena x2 adalah salinan dari nilai x
	// sehingga perubahan pada x2 tidak mempengaruhi nilai x
	// jika ingin mengubah nilai x melalui x2, maka harus menggunakan pointer
	fmt.Println(x)
	fmt.Println(x2)
}

func parameterPointer(ori *int, val int) {
	*ori = val
}

func main() {
	// pointerExample()
	// withoutPointer()

	var number int = 10
	fmt.Println("Before:", number) // Before: 10
	parameterPointer(&number, 25)
	fmt.Println("After:", number) // After: 25
}
