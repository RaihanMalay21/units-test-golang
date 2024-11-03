package helper

import (
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/assert"
	"runtime"
	"testing"
	"fmt"
)

// menjalankan semua bechmarck sekaligus dengan unit test
// go test -v bench=
// menjalankan bechmark semua bechmarck tanpa unit test
// go test -v -run=NotMathUnitTest -bench=.
// menjalankan bechmark tertentu
// go test -v -run=NotMathUnitTest -bech=nameBechMarck
// menjalankan unit test di root module
// go test -v -bench=. ./..
// hilangkan -run ketika di run
func BenchmarkSub(b *testing.B) {
	b.Run("Raihan", func(b *testing.B) {
		for i:=0; i < b.N; i++ {
			HelloWorld("Raihan")
		}
	})
	b.Run("Malay", func(b *testing.B) {
		for i:=0; i < b.N; i++ {
			HelloWorld("Malay")
		}
	})
}

// func BechMarckHelloWorld(b *testing.B) {
// 	for i:=0; i < b.N; i++ {
// 		HelloWorld("Raihan")
// 	}
// }

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name string
		request string
		expected string
	}{
		{
			name: "HelloWorld(Raihan)",
			request: "Raihan",
			expected: "Hello Raihan",
		},
		{
			name: "HelloWorld(Malay)",
			request: "Malay",
			expected: "Hello Malay",
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
	t.Run("Raihan", func(t *testing.T) {
		result := HelloWorld("Raihan")
		assert.Equal(t, "Hello Raihan", result, "result must be Hello Raihan")
	})
	t.Run("Malay", func(t *testing.T) {
		result := HelloWorld("Malay")
		assert.Equal(t, "Hello Malay", result, "result must be Hello Malay")
	})
}

func TestMain(m *testing.M) {
	fmt.Println("Before Unit Test")

	m.Run()

	fmt.Println("After Unit Test")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Program cannot run on windows")
	} 

	result := HelloWorld("Raihan")
	assert.Equal(t, "Hello Raihan", result)
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Raihan")
	require.Equal(t, "Hello Raiha", result, "result have be Hello Raiha")
	fmt.Println("hehehehehehehhhhhheeeee!")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Raihan")
	assert.Equal(t, "Hello Raihan", result, "result must be Hello Raihan")
	assert.NotEqual(t, "Hello Raihan", result, "result must not be Hello Raihan")
	fmt.Println("Heeee")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Raihan")
	if result != "Hello Raihan" {
		t.Error("Result is not Hello Raihan")
		// t.Fail()
		// panic("Result is not Hello Raihan")
	}
	fmt.Println("Hello world Raihan")
}

func TestHelloWorldMalay(t *testing.T) {
	result := HelloWorld("Malay")
	if result != "Hello Malay" {
		t.Fatal("Result is not Hello Malay")
		// t.FailNow()
		// panic("Result is not Hello Raihan")
	}
	fmt.Println("Hello world Malay")
}
