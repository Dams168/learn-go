package main

/*
   Package strconv di Go
   Digunakan untuk konversi tipe data dari string ke tipe data lain dan sebaliknya.

*/

import (
	"fmt"
	"strconv"
)

func main() {
	result, err := strconv.Atoi("1234")
	if err != nil {
		fmt.Println("Error converting string to int:", err)
	}

	fmt.Println("Converted integer:", result)

	var stringValue string = strconv.Itoa(5678)
	fmt.Println("Converted string:", stringValue)

	binary := strconv.FormatInt(505, 2)
	fmt.Println("Binary representation of 505 is", binary)
}
