package iteration

import "testing"

func TestRepeat(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, repeated, expected string) {
		t.Helper()
		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	}

	t.Run("repeat without count", func(t *testing.T) {
		repeated := Repeat("a", "", 0)
		expected := "aaaaa"
		assertCorrectMessage(t, repeated, expected)
	})

	t.Run("repeat with count", func(t *testing.T) {
		repeated := Repeat("a", "", 5)
		expected := "aaaaa"
		assertCorrectMessage(t, repeated, expected)
	})

	t.Run("repeat with count and replace character", func(t *testing.T) {
		repeated := Repeat("a", "b", 5)
		expected := "bbbbb"
		assertCorrectMessage(t, repeated, expected)
	})

	t.Run("remove space characters", func(t *testing.T) {
		repeated := Trim("a a a a a")
		expected := "aaaaa"
		assertCorrectMessage(t, repeated, expected)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", "", 0)
	}
}
