package unittest

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkHelloWorldTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Dam",
			request: "Dam",
		},
		{
			name:    "Dim",
			request: "Dim",
		},
	}
	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

func BenchmarkHelloWorldSub(b *testing.B) {
	b.Run("Dam", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Dam")
		}
	})
	b.Run("Dim", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Dim")
		}
	})
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Dam")
	}
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Test for Dam",
			request:  "Dam",
			expected: "Hello Dam",
		},
		{
			name:     "Test for Dim",
			request:  "Dim",
			expected: "Hello Dim",
		},
		{
			name:     "Test for John",
			request:  "John",
			expected: "Hello John",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			assert.Equal(t, test.expected, result, "Result must be "+test.expected)
		})
	}
}

func TestSubTestHelloWorld(t *testing.T) {
	t.Run("Dam", func(t *testing.T) {
		result := HelloWorld("Dam")
		assert.Equal(t, "Hello Dam", result, "Result must be 'Hello Dam'")
	})
	t.Run("Dim", func(t *testing.T) {
		result := HelloWorld("Dim")
		assert.Equal(t, "Hello Dim", result, "Result must be 'Hello Dim'")
	})
}

func TestMain(m *testing.M) {
	fmt.Println("Before Unit Test")
	m.Run()
	fmt.Println("After Unit Test")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Skipping test on Windows OS")
	}

	result := HelloWorld("Dam")
	assert.Equal(t, "Hello Dam", result, "Result must be 'Hello Dam'")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Dam")

	assert.Equal(t, "Hello Dam", result, "Result must be 'Hello Dam'")
}

func TestHelloWorldDam(t *testing.T) {
	result := HelloWorld("Dam")

	if result != "Hello Dam" {
		t.Errorf("Expected 'Hello Dam', but got '%s'", result)
	}
}

func TestHelloWorldJohn(t *testing.T) {
	result := HelloWorld("John")
	if result != "Hello John" {
		t.Errorf("Expected 'Hello John', but got '%s'", result)
	}
}
