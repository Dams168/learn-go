package generics

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Data[T any] artinya:
// Struct ini punya type parameter T.
// T bisa tipe apa saja (any).
//
// Jadi Data bisa menjadi:
//
//	Data[string]
//	Data[int]
//	Data[float64]
//	Data[User]
//
// dll.
type Data[T any] struct {
	First  T
	Second T
}

// Perhatikan ini:
// (d *Data[_])
//
// _ artinya kita tidak peduli dengan tipe T.
// Method ini tidak menggunakan T sama sekali.
//
// Jadi method ini berlaku untuk SEMUA Data[T],
// tanpa peduli apa T-nya.
func (d *Data[_]) SayHello(name string) string {
	return "Hello " + name
}

// Di sini kita pakai T.
// Artinya method ini tergantung pada type parameter T.
//
// Kalau Data[string] → T = string
// Kalau Data[int] → T = int
func (d *Data[T]) ChangeFirst(first T) T {
	d.First = first
	return d.First
}

func TestData(t *testing.T) {

	// Membuat Data dengan tipe string
	data := Data[string]{
		First:  "Eko",
		Second: "Khannedy",
	}

	fmt.Println(data)
}

func TestGenericMethod(t *testing.T) {

	data := Data[string]{
		First:  "Eko",
		Second: "Khannedy",
	}

	// Karena Data[string],
	// maka ChangeFirst hanya menerima string.
	assert.Equal(t, "Budi", data.ChangeFirst("Budi"))

	// Method ini tidak tergantung T.
	assert.Equal(t, "Hello Eko", data.SayHello("Eko"))
}
