package main

import (
	"fmt"
	"strings"
)

func printMessage(msg string, arr []string) {
	print := strings.Join(arr, " ")

	fmt.Println(msg, print)
}

func PersegiPanjang(p, l int) (int, int) {

	var keliling = 2 * (p + l)
	var luas = p * l

	return keliling, luas
}

// fungsi variadic
func averageNum(nums ...int) float64 {
	var total = 0

	for _, n := range nums {
		total += n
	}

	avg := float64(total) / float64(len(nums))

	return avg
}

func introduce(name string, hobbies ...string) {
	var hobbyString = strings.Join(hobbies, ", ")

	fmt.Println("Hello, My Name : ", name)
	fmt.Println("My Hobbies are : ", hobbyString)
}

// fungsi closure => fungsi yang disimpan divariabel, digunakan untuk membungkus sebuah fungsi yang hanya dijalankan sekali saja / sekali pakai
func counter() func() int {
	count := 0
	increment := func() int {
		count++
		return count
	}
	return increment
}
