package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// func TestHelloWorld(t *testing.T) {
// 	result := HelloWorld("budi")
// 	if result != "hello budi" {
// 		panic("result is not hello budi")
// 	}
// }

func TestMain(m *testing.M) {
	fmt.Println("sebelum")
	m.Run()
	fmt.Println("sesudah")
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("budi")

	}
}
func BenchmarkHelloWorldBudi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("budi")

	}
}

func TestTableTest(t *testing.T) {

	data := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "budi",
			request:  "budi",
			expected: "hi budi",
		},
		{
			name:     "tono",
			request:  "tono",
			expected: "hi tono",
		},
		{
			name:     "rangga",
			request:  "rangga",
			expected: "hi rangga",
		},
	}

	for _, item := range data {
		t.Run(item.name, func(t *testing.T) {
			result := HelloWorld((item.request))
			assert.Equal(t, item.expected, result)

		})
	}

}

func TestSubTest(t *testing.T) {

	t.Run("budi", func(t *testing.T) {
		result := HelloWorld("budi")
		require.Equal(t, "hello budi", result, "result must be 'hi budi'")
	})

	t.Run("taryo", func(t *testing.T) {
		result := HelloWorld("taryo")
		require.Equal(t, "hello taryo", result, "result must be 'hi taryo'")
	})

}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("tidak bisa berjalan di mac os")
	}
	result := HelloWorld("budi")
	assert.Equal(t, "hi budi", result, "result must be 'hi budi'")
	fmt.Println("TestHelloworld with skip and assert done")
}

func TestHelloWorldFail(t *testing.T) {
	result := HelloWorld("budi")
	if result != "hello budi" {
		t.Fail()
	}
	fmt.Println("sudah error tapi tetap lanjut")
}

func TestHelloWorldFailNow(t *testing.T) {
	result := HelloWorld("budi")
	if result != "hello budi" {
		t.FailNow()
	}
	fmt.Println("sudah error tapi tidak lanjut")
}

func TestHelloWorldError(t *testing.T) {
	result := HelloWorld("budi")
	if result != "hello budi" {
		t.Error("harusnya hi budi")
	}
	fmt.Println("sudah error tapi tetap lanjut saja")
}

func TestHelloWorldFatal(t *testing.T) {
	result := HelloWorld("budi")
	if result != "hello budi" {
		t.Fatal("harusnya hi budi")
	}
	fmt.Println("sudah error tapi tidak lanjut saja")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("budi")
	require.Equal(t, "hello budi", result, "result must be 'hi budi'")
	fmt.Println("TestHelloworld with require done")
}
func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("budi")
	assert.Equal(t, "hello budi", result, "result must be 'hi budi'")
	fmt.Println("TestHelloworld with assert done")
}
