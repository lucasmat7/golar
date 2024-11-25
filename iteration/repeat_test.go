package iteration

import (
	"fmt"
	"testing"
)

// tests

func TestRepeat(t *testing.T) {
	t.Run("testing dummy 5a's", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"

		assertTests(t, repeated, expected)
	})

	t.Run("testing no repeats", func(t *testing.T) {
		repeated := Repeat("y", 0)
		expected := ""

		assertTests(t, repeated, expected)
	})
}

// helper functions

func assertTests(t testing.TB, repeated, expected string) {
	t.Helper()
	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

// examples

func ExampleRepeat() {
	repeated := Repeat("a", 5)
	fmt.Printf(repeated)
	// Output: aaaaa
}

// benchmarks

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
