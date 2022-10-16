package iteration

import (
	"testing"
)

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func assertCorrectOutput(t *testing.T, received, expected string) {
	t.Helper()

	if received != expected {
		t.Errorf("expected %q but got %q", expected, received)
	}
}

func TestRepeat(t *testing.T) {
	var received, expected string

	t.Run("With a positive repeat value", func(t *testing.T) {
		received = Repeat("b", 4)
		expected = "bbbb"

		assertCorrectOutput(t, received, expected)
	})

	t.Run("With repeat value as 0", func(t *testing.T) {
		received = Repeat("c", 0)
		expected = ""

		assertCorrectOutput(t, received, expected)
	})

	t.Run("With a negative repeat value", func(t *testing.T) {
		received = Repeat("d", -5)
		expected = ""

		assertCorrectOutput(t, received, expected)
	})
}
