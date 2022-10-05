package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkTable(b *testing.B) {
	benchs := []struct {
		name    string
		request string
	}{
		{
			name:    "HelloWorld(Rakha)",
			request: "Rakha",
		},
		{
			name:    "HelloWorld(Eko)",
			request: "Eko",
		},
	}

	for _, bench := range benchs {
		b.Run(bench.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(bench.request)
			}
		})
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Rakha", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Rakha")
		}
	})

	b.Run("Eko", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Eko")
		}
	})
}

func Benchmark_Hello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Eko")
	}
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Rakha")
	}
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "HelloWorld(Rakha)",
			request:  "Rakha",
			expected: "Hello Rakha",
		},
		{
			name:     "HelloWorld(Eko)",
			request:  "Eko",
			expected: "Hello Eko",
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
	//run Rakha
	t.Run("Rakha", func(t *testing.T) {
		result := HelloWorld("Rakha")
		require.Equal(t, "Hello Rakha", result)

	})

	//run Eko
	t.Run("Eko", func(t *testing.T) {
		result := HelloWorld("Eko")
		require.Equal(t, "Hello Eko", result)

	})

}

func TestMain(m *testing.M) {
	//before unit test
	fmt.Println("Before Unit Test")

	//eksekusi semua unit test
	m.Run()

	//after unit test
	fmt.Println("After Unit Test")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Rakha")

	if result != "Hello Rakha" {
		//err
		panic("result is not 'Hello Rakha'")
	}
}

func TestHelloWorldAssertion(t *testing.T) {
	result := HelloWorld("Rakha")

	assert.Equal(t, "Hello Rakha", result)

}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Rakha")

	require.Equal(t, "Hello Rakha", result)
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("cannot run on mac")
	}

	result := HelloWorld("Rakha")
	require.Equal(t, "Hello Rakha", result)
}
