package generics

import "testing"

// untuk menambahkan lebih dari satu type parameter, kita dapat menggunakan tanda koma untuk memisahkan setiap type parameter yang ingin kita gunakan.
func MultipleTypeParameter[T1 any, T2 any](param1 T1, param2 T2) (T1, T2) {
	return param1, param2
}

func TestMultipleTypeParameter(t *testing.T) {
	MultipleTypeParameter[int, string](100, "test")
	MultipleTypeParameter[string, int]("test", 100)
}
