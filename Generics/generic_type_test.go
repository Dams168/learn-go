package generics

import (
	"fmt"
	"testing"
)

// Bag[T any] adalah generic type.
// Artinya Bag bisa menyimpan tipe data apa saja.
//
// Ini sebenarnya alias dari slice:
//
//	[]T
//
// Jadi:
//
//	Bag[string] = []string
//	Bag[int]    = []int
type Bag[T any] []T

// PrintBag menerima parameter Bag[T].
//
// T any artinya:
// T bisa tipe apa saja.
func PrintBag[T any](bag Bag[T]) {
	for _, value := range bag {
		fmt.Println(value)
	}
}

func TestBagString(t *testing.T) {

	// Membuat Bag dengan tipe string
	// Sama seperti []string
	names := Bag[string]{"Eko", "Kurniawan", "Khannedy"}

	PrintBag(names)
}

func TestBagInt(t *testing.T) {

	// Membuat Bag dengan tipe int
	numbers := Bag[int]{1, 2, 3, 4, 5}

	// Bisa eksplisit tulis [int]
	PrintBag[int](numbers)
}
