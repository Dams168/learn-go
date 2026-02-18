package generics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Age adalah custom type dengan underlying type int.
// Artinya Age BUKAN int, tapi berbasis int.
type Age int

// Number adalah interface constraint khusus untuk generics.
//
// Tanda | artinya "atau".
//
// ~int artinya:
// Semua tipe yang underlying type-nya int.
// Termasuk:
//   - int
//   - type Age int
//
// Tanpa ~, custom type seperti Age tidak akan cocok.
type Number interface {
	~int | int8 | int16 | int32 | int64 |
		float32 | float64
}

// T Number artinya:
// T hanya boleh tipe yang termasuk constraint Number.
//
// Karena semua tipe Number mendukung operator <,
// maka kita bisa pakai perbandingan.
func Min[T Number](first, second T) T {
	if first < second {
		return first
	} else {
		return second
	}
}

func TestMin(t *testing.T) {

	// Secara eksplisit menentukan tipe generic
	assert.Equal(t, 100, Min[int](100, 200))

	assert.Equal(t, int64(100), Min[int64](int64(100), int64(200)))

	assert.Equal(t, float64(100), Min[float64](float64(100), float64(200)))

	// Age bisa dipakai karena ~int
	assert.Equal(t, Age(100), Min[Age](Age(100), Age(200)))
}

func TestMinTypeInference(t *testing.T) {

	// Go bisa menebak tipe T otomatis
	assert.Equal(t, 100, Min(100, 200))

	assert.Equal(t, int64(100), Min(int64(100), int64(200)))

	assert.Equal(t, float64(100), Min(float64(100), float64(200)))

	assert.Equal(t, Age(100), Min(Age(100), Age(200)))
}
