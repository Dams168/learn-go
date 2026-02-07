package main

/*
   Package math di Go
   Digunakan untuk melakukan operasi matematika dasar dan lanjutan.
*/

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Ceil(1.45))
	fmt.Println(math.Round(1.45))
	fmt.Println(math.Floor(1.45))
	fmt.Println(math.Sqrt(16))
	fmt.Println(math.Pow(2, 3))
	fmt.Println(math.Min(2, 3))
}
