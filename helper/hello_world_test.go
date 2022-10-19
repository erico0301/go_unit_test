package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Eko")
	}
}
func BenchmarkHelloWorldErico(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Erico")
	}
}

func TestMain(m *testing.M) {
	// before
	fmt.Println("BEFORE UNIT TEST")

	// Process
	m.Run()

	// After
	fmt.Println("AFTER UNIT TEST")
}

// Table Test
func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Eko",
			request:  "Eko",
			expected: "Hello Eko",
		}, {
			name:     "Erico",
			request:  "Erico",
			expected: "Hello Erico",
		}, {
			name:     "Ericq",
			request:  "Ericq",
			expected: "Hello Ericq",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			assert.Equal(t, test.expected, result)
		})
	}

}

func TestSubTest(t *testing.T) {
	t.Run("Eko", func(t *testing.T) {
		result := HelloWorld("Eko")
		assert.Equal(t, "Hello Eko", result, "Result must be 'Hello Eko'")
	})
	t.Run("Erico", func(t *testing.T) {
		result := HelloWorld("Erico")
		assert.Equal(t, "Hello Erico", result, "Result must be 'Hello Eko'")
	})
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Can not run on Windows")
	}
	result := HelloWorld("Eko")
	assert.Equal(t, "Hello Eko", result, "Result must be 'Hello Eko'")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Eko")
	assert.Equal(t, "Hello Eko", result, "Result must be 'Hello Eko'")
	fmt.Println("TestHelloWorld with Assert Done")
}
func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Eko")
	require.Equal(t, "Hello Eko", result, "Result must be 'Hello Eko'")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Eko")
	if result != "Hello Eko" {
		// Unit test failed
		t.Error("Result Not Hello Eko")
	}
	// fmt.Println("TestHelloWorldEko Done")
}

func TestHelloWorldErico(t *testing.T) {
	result := HelloWorld("Erico")
	if result != "Hello Erico" {
		// Unit test failed
		t.Error("Result Not Hello Erico")
	}
	// fmt.Println("TestHelloWorldErico Done")
}
