package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("Should return the given string repeat the given number of repetitions", func(t *testing.T) {
		repeated := Repeat("a", 10)
		expected := "aaaaaaaaaa"

		if repeated != expected {
			t.Errorf("Expected %q but got %q", expected, repeated)
		}
	})

	t.Run("Should return the given string repeated 5 times(default) if repetitions argument is <= 0", func(t *testing.T) {
		resultZero := Repeat("z", 0)
		expectedZero := "zzzzz"

		if resultZero != expectedZero {
			t.Errorf("Expected %q but got %q", expectedZero, resultZero)
		}

		resultLessZero := Repeat("y", -5)
		expectedLessZero := "yyyyy"

		if resultLessZero != expectedLessZero {
			t.Errorf("Expected %q but got %q", expectedLessZero, resultLessZero)
		}
	})

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}
