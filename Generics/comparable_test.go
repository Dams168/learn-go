package generics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// comparable adalah constraint yang digunakan untuk membatasi type parameter hanya pada tipe data yang dapat dibandingkan menggunakan operator perbandingan seperti == dan !=. Dengan menggunakan comparable,
// kita dapat memastikan bahwa hanya tipe data yang dapat dibandingkan yang dapat digunakan dengan fungsi, tipe data, atau metode tertentu.
func ComparableTypeParameter[T comparable](param1, param2 T) bool {
	if param1 == param2 {
		return true
	} else {
		return false
	}
}

func TestComparableTypeParameter(t *testing.T) {
	assert.True(t, ComparableTypeParameter[int](100, 100))
	assert.False(t, ComparableTypeParameter[string]("test", "TEST"))
}
