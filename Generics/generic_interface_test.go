package generics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// GetterSetter[T any] adalah interface generic.
// Artinya interface ini fleksibel terhadap tipe T.
//
// Jika T = string → GetValue() string
// Jika T = int    → GetValue() int
type GetterSetter[T any] interface {
	GetValue() T
	SetValue(value T)
}

// Function ini menerima parameter
// yang HARUS mengimplementasikan
// GetterSetter[T]
func ChangeValue[T any](param GetterSetter[T], value T) T {
	param.SetValue(value)
	return param.GetValue()
}

type MyData[T any] struct {
	Value T
}

// Method ini membuat MyData[T]
// memenuhi interface GetterSetter[T]
func (d *MyData[T]) GetValue() T {
	return d.Value
}

func (d *MyData[T]) SetValue(value T) {
	d.Value = value
}

func TestGenericInterface(t *testing.T) {
	myData := MyData[string]{}

	result := ChangeValue[string](&myData, "Eko")

	assert.Equal(t, "Eko", result)
	assert.Equal(t, "Eko", myData.Value)
}
