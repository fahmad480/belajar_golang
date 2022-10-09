package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Faraaz",
			request: "Faraaz",
		},
		{
			name:    "Ahmad",
			request: "Ahmad",
		},
		{
			name:    "FaraazAhmadPermadi",
			request: "Faraaz Ahmad Permadi",
		},
		{
			name:    "Budi",
			request: "Budi Nugraha",
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

func BenchmarkSub(b *testing.B) {
	b.Run("Faraaz", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Faraaz")
		}
	})
	b.Run("Ahmad", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Ahmad")
		}
	})
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Faraaz")
	}
}

func BenchmarkHelloWorldAhmad(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Ahmad")
	}
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Faraaz",
			request:  "Faraaz",
			expected: "Hello Faraaz",
		},
		{
			name:     "Ahmad",
			request:  "Ahmad",
			expected: "Hello Ahmad",
		},
		{
			name:     "Permadi",
			request:  "Permadi",
			expected: "Hello Permadi",
		},
		{
			name:     "Budi",
			request:  "Budi",
			expected: "Hello Budi",
		},
		{
			name:     "Joko",
			request:  "Joko",
			expected: "Hello Joko",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Faraaz", func(t *testing.T) {
		result := HelloWorld("Faraaz")
		require.Equal(t, "Hello Faraaz", result, "Result must be 'Hello Faraaz'")
	})
	t.Run("Ahmad", func(t *testing.T) {
		result := HelloWorld("Ahmad")
		require.Equal(t, "Hello Ahmad", result, "Result must be 'Hello Ahmad'")
	})
	t.Run("Permadi", func(t *testing.T) {
		result := HelloWorld("Permadi")
		require.Equal(t, "Hello Permadi", result, "Result must be 'Hello Permadi'")
	})
}

func TestMain(m *testing.M) {
	// before
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	// after
	fmt.Println("AFTER UNIT TEST")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Can not run on Mac OS")
	}

	result := HelloWorld("Faraaz")
	require.Equal(t, "Hello Faraaz", result, "Result must be 'Hello Faraaz'")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Faraaz")
	require.Equal(t, "Hello Faraaz", result, "Result must be 'Hello Faraaz'")
	fmt.Println("TestHelloWorld with Require Done")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Faraaz")
	assert.Equal(t, "Hello Faraaz", result, "Result must be 'Hello Faraaz'")
	fmt.Println("TestHelloWorld with Assert Done")
}

func TestHelloWorldFaraaz(t *testing.T) {
	result := HelloWorld("Faraaz")

	if result != "Hello Faraaz" {
		// error
		t.Error("Result must be 'Hello Faraaz'")
	}

	fmt.Println("TestHelloWorldFaraaz Done")
}

func TestHelloWorldPermadi(t *testing.T) {
	result := HelloWorld("Permadi")

	if result != "Hello Permadi" {
		// error
		t.Fatal("Result must be 'Hello Permadi'")
	}

	fmt.Println("TestHelloWorldPermadi Done")
}

func TestHelloWorldAhmad(t *testing.T) {
	result := HelloWorld("Ahmad")

	if result != "Hello Ahmad" {
		// error
		t.Fatal("Result must be 'Hello Ahmad'")
	}

	fmt.Println("TestHelloWorldAhmad Done")
}
