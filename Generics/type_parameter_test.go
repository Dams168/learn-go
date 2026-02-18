package generics

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// type any adalah alias untuk interface{} yang dapat menerima nilai dari tipe data apa pun.
func TypeParameter[T any](param T) T {
	fmt.Println(param)
	return param
}

func TestTypeParameter(t *testing.T) {
	var res string = TypeParameter[string]("test")
	assert.Equal(t, "test", res)

	var res2 int = TypeParameter[int](100)
	assert.Equal(t, 100, res2)
}
