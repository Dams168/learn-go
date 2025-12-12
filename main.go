package main

import "fmt"

// package i/o

func main() {

	// Variable & Type Data

	/*

		var <nama-variabel> <tipe-data>
		var <nama-variabel> <tipe-data> = <nilai>

		:= <nilai>

		var firstName string = "john"
		lastName := "Doe"

		one, isFriday, twoPointTwo, say := 1, true, 2.2, "hello"
		_ =  "wick"



		const (
			pi = 3.14
			e  = 2.71
			isToday bool = false
			numeric int8 = 42
			num
		)

		fmt.Println("Hello", firstName, lastName + "!","!")
		fmt.Println(one, isFriday, twoPointTwo, say)
		fmt.Println(pi, e, isToday, numeric, num)
	*/

	// array
	// var names [4]string
	// names[0] = "trafalgar"
	// names[1] = "d"
	// names[2] = "water"
	// names[3] = "law"

	// fmt.Println(names[0], names[1], names[2], names[3])

	// var fruits = [4]string{"apple", "grape", "banana", "melon"}
	// fmt.Println("Jumlah element \t\t", len(fruits))
	// fmt.Println("Isi semua element \t", fruits)

	// condition(8)
	// looping()
	// FizzBuzz(20)
	// factorial(5)

	// arrayExample()
	// sliceExample()
	mapExample()

}

func FizzBuzz(point int) {
	for i := 1; i <= point; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func condition(nilai int) {

	// conditional statement
	if nilai == 10 {
		fmt.Println("Perfect")
	} else if nilai >= 7 && nilai < 10 {
		fmt.Println("Great")
	} else if nilai >= 5 && nilai < 7 {
		fmt.Println("Good")
	} else {
		fmt.Println("Nice Try")
	}

	// conditional statement with temporary variable
	if grade := nilai % 2; grade == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	// switch case
	switch nilai {
	case 10:
		fmt.Println("Perfect")
	case 7, 8, 9:
		fmt.Println("Great")
	case 5, 6:
		fmt.Println("Good")
	default:
		fmt.Println("Nice Try")
	}

}

func looping() {
	//for
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// while
	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}

	// loop wothout argue

	// for {
	// 	fmt.Println(j)
	// 	j++
	// 	if j == 10 {
	// 		break
	// 	}
	// }

	// loop woth range for array, string, slice, map
	var x string = "yahaha"
	for i, v := range x {
		fmt.Println("indeks = ", i, " Value = ", v)
	}

}
